import {assert} from 'chai';
import Outpoint from './Outpoint';
import PlasmaContract from '../lib/PlasmaContract';
import PlasmaClient from '../lib/PlasmaClient';
import MerkleTree from '../lib/MerkleTree';
import {sha256} from '../lib/hash';
import {ethSign} from '../lib/sign';
import BN = require('bn.js');

export default class ExitOperation {
  private contract: PlasmaContract;

  private client: PlasmaClient;

  private from: string;

  private outpoint: Outpoint | null = null;

  private committedFee: BN | null = null;

  constructor (contract: PlasmaContract, client: PlasmaClient, from: string) {
    this.contract = contract;
    this.client = client;
    this.from = from;
  }

  withOutpoint (outpoint: Outpoint): ExitOperation {
    this.outpoint = outpoint;
    return this;
  }

  withCommittedFee (committedFee: BN): ExitOperation {
    this.committedFee = committedFee;
    return this;
  }

  async exit (privateKey: Buffer) {
    assert(this.outpoint, 'an outpoint to exit must be provided');
    assert(this.committedFee, 'a fee must be provided');
    const block = await this.client.getBlock(this.outpoint!.blockNum);
    const merkle = new MerkleTree();
    for (const tx of block.transactions) {
      merkle.addItem(sha256(tx.toRLP()));
    }
    const {proof} = merkle.generateProofAndRoot(this.outpoint!.txIdx);
    const confirmedTx = this.outpoint!.transaction;
    const callSigHash = sha256(Buffer.concat([sha256(confirmedTx.toRLP()), block.header.merkleRoot]));
    const callSig = ethSign(callSigHash, privateKey);
    const callSigs: [Buffer, Buffer] = confirmedTx.input1.owner === this.from ? [callSig, callSig] : [callSig, Buffer.from('')];
    await this.contract.startExit(this.outpoint!, proof, callSigs, this.committedFee!, this.from);
  }
}