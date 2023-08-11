import Transaction from './Transaction';
import {sha256} from '../crypto/hash';
import {ethSign} from '../crypto/sign';

/**
 * Wrapper class for a Plasma transaction and its associated `confirmSignatures`.
 * For transactions that haven't been confirmed yet, each element of the
 * `confirmSignatures` array will be a zero-filled `Buffer` of length 65.
 */
export default class ConfirmedTransaction {
  /**
   * The transaction itself.
   */
  public readonly transaction: Transaction;

  /**
   * The list of confirm signatures for this transaction.
   */
  public confirmSignatures: [Buffer, Buffer] | null;

  constructor (transaction: Transaction, confirmSignatures: [Buffer, Buffer] | null) {
    this.transaction = transaction;
    this.confirmSignatures = confirmSignatures;
  }

  /**
   * Returns the hash of the transaction used to generate the
   * transaction's `confirmSignatures`. According to the specification,
   * `confirmSignatures` are generated by taking the `sha256 hash of
   * the concatenation of:
   *
   * 1. The `sha256` hash of the RLP-encoded transaction.
   * 2. The `merkleRoot` of the block that includes this transaction.
   *
   * This method will throw an error if the wrapped transaction does not have
   * a `blockNum` set. A transaction without a `blockNum` set has not been
   * included in a block, and so cannot be confirmed.
   *
   * @param merkleRoot The `merkleRoot` of the block that includes this transaction.
   */
  confirmHash (merkleRoot: Buffer) {
    if (this.transaction.body.blockNum === 0) {
      throw new Error('refusing to generate confirmHash for an unconfirmed transaction.');
    }

    return sha256(Buffer.concat([
      sha256(this.transaction.toRLP()),
      merkleRoot,
    ]));
  }

  /**
   * Sets this `ConfirmedTransaction`'s confirmation signatures by
   * signing the result of `confirmHash` with the provided `privateKey`.
   *
   * Note that this method sets the `confirmSignatures` on the class itself
   * and does not return them.
   *
   * This method will throw an error if the wrapped transaction does not have
   * a `blockNum` set. A transaction without a `blockNum` set has not been
   * included in a block, and so cannot be confirmed.
   *
   * @param privateKey The private key to create the `confirmSignatures` with.
   * @param merkleRoot The `merkleRoot` of the block that includes this transaction.
   */
  confirmSign (privateKey: Buffer, merkleRoot: Buffer) {
    const confirmSigHash = this.confirmHash(merkleRoot);
    const confirmSig = ethSign(confirmSigHash, privateKey);
    this.confirmSignatures = [
      confirmSig, confirmSig,
    ];
  }
}