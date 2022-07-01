/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "vigorousdeveloper.pochuman.pochuman";

export interface TransactionData {
  index: string;
  originChain: string;
  originAddress: string;
  targetChain: string;
  targetAddress: string;
  amount: string;
  time: string;
  creator: string;
  status: string;
  confirmedBlockHash: string;
  signedKey: string;
  fee: string;
}

const baseTransactionData: object = {
  index: "",
  originChain: "",
  originAddress: "",
  targetChain: "",
  targetAddress: "",
  amount: "",
  time: "",
  creator: "",
  status: "",
  confirmedBlockHash: "",
  signedKey: "",
  fee: "",
};

export const TransactionData = {
  encode(message: TransactionData, writer: Writer = Writer.create()): Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    if (message.originChain !== "") {
      writer.uint32(18).string(message.originChain);
    }
    if (message.originAddress !== "") {
      writer.uint32(26).string(message.originAddress);
    }
    if (message.targetChain !== "") {
      writer.uint32(34).string(message.targetChain);
    }
    if (message.targetAddress !== "") {
      writer.uint32(42).string(message.targetAddress);
    }
    if (message.amount !== "") {
      writer.uint32(50).string(message.amount);
    }
    if (message.time !== "") {
      writer.uint32(58).string(message.time);
    }
    if (message.creator !== "") {
      writer.uint32(66).string(message.creator);
    }
    if (message.status !== "") {
      writer.uint32(74).string(message.status);
    }
    if (message.confirmedBlockHash !== "") {
      writer.uint32(82).string(message.confirmedBlockHash);
    }
    if (message.signedKey !== "") {
      writer.uint32(90).string(message.signedKey);
    }
    if (message.fee !== "") {
      writer.uint32(98).string(message.fee);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): TransactionData {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseTransactionData } as TransactionData;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        case 2:
          message.originChain = reader.string();
          break;
        case 3:
          message.originAddress = reader.string();
          break;
        case 4:
          message.targetChain = reader.string();
          break;
        case 5:
          message.targetAddress = reader.string();
          break;
        case 6:
          message.amount = reader.string();
          break;
        case 7:
          message.time = reader.string();
          break;
        case 8:
          message.creator = reader.string();
          break;
        case 9:
          message.status = reader.string();
          break;
        case 10:
          message.confirmedBlockHash = reader.string();
          break;
        case 11:
          message.signedKey = reader.string();
          break;
        case 12:
          message.fee = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): TransactionData {
    const message = { ...baseTransactionData } as TransactionData;
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    if (object.originChain !== undefined && object.originChain !== null) {
      message.originChain = String(object.originChain);
    } else {
      message.originChain = "";
    }
    if (object.originAddress !== undefined && object.originAddress !== null) {
      message.originAddress = String(object.originAddress);
    } else {
      message.originAddress = "";
    }
    if (object.targetChain !== undefined && object.targetChain !== null) {
      message.targetChain = String(object.targetChain);
    } else {
      message.targetChain = "";
    }
    if (object.targetAddress !== undefined && object.targetAddress !== null) {
      message.targetAddress = String(object.targetAddress);
    } else {
      message.targetAddress = "";
    }
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = String(object.amount);
    } else {
      message.amount = "";
    }
    if (object.time !== undefined && object.time !== null) {
      message.time = String(object.time);
    } else {
      message.time = "";
    }
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.status !== undefined && object.status !== null) {
      message.status = String(object.status);
    } else {
      message.status = "";
    }
    if (
      object.confirmedBlockHash !== undefined &&
      object.confirmedBlockHash !== null
    ) {
      message.confirmedBlockHash = String(object.confirmedBlockHash);
    } else {
      message.confirmedBlockHash = "";
    }
    if (object.signedKey !== undefined && object.signedKey !== null) {
      message.signedKey = String(object.signedKey);
    } else {
      message.signedKey = "";
    }
    if (object.fee !== undefined && object.fee !== null) {
      message.fee = String(object.fee);
    } else {
      message.fee = "";
    }
    return message;
  },

  toJSON(message: TransactionData): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    message.originChain !== undefined &&
      (obj.originChain = message.originChain);
    message.originAddress !== undefined &&
      (obj.originAddress = message.originAddress);
    message.targetChain !== undefined &&
      (obj.targetChain = message.targetChain);
    message.targetAddress !== undefined &&
      (obj.targetAddress = message.targetAddress);
    message.amount !== undefined && (obj.amount = message.amount);
    message.time !== undefined && (obj.time = message.time);
    message.creator !== undefined && (obj.creator = message.creator);
    message.status !== undefined && (obj.status = message.status);
    message.confirmedBlockHash !== undefined &&
      (obj.confirmedBlockHash = message.confirmedBlockHash);
    message.signedKey !== undefined && (obj.signedKey = message.signedKey);
    message.fee !== undefined && (obj.fee = message.fee);
    return obj;
  },

  fromPartial(object: DeepPartial<TransactionData>): TransactionData {
    const message = { ...baseTransactionData } as TransactionData;
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    if (object.originChain !== undefined && object.originChain !== null) {
      message.originChain = object.originChain;
    } else {
      message.originChain = "";
    }
    if (object.originAddress !== undefined && object.originAddress !== null) {
      message.originAddress = object.originAddress;
    } else {
      message.originAddress = "";
    }
    if (object.targetChain !== undefined && object.targetChain !== null) {
      message.targetChain = object.targetChain;
    } else {
      message.targetChain = "";
    }
    if (object.targetAddress !== undefined && object.targetAddress !== null) {
      message.targetAddress = object.targetAddress;
    } else {
      message.targetAddress = "";
    }
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = object.amount;
    } else {
      message.amount = "";
    }
    if (object.time !== undefined && object.time !== null) {
      message.time = object.time;
    } else {
      message.time = "";
    }
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.status !== undefined && object.status !== null) {
      message.status = object.status;
    } else {
      message.status = "";
    }
    if (
      object.confirmedBlockHash !== undefined &&
      object.confirmedBlockHash !== null
    ) {
      message.confirmedBlockHash = object.confirmedBlockHash;
    } else {
      message.confirmedBlockHash = "";
    }
    if (object.signedKey !== undefined && object.signedKey !== null) {
      message.signedKey = object.signedKey;
    } else {
      message.signedKey = "";
    }
    if (object.fee !== undefined && object.fee !== null) {
      message.fee = object.fee;
    } else {
      message.fee = "";
    }
    return message;
  },
};

type Builtin = Date | Function | Uint8Array | string | number | undefined;
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;
