/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "vigorousdeveloper.pochuman.pochuman";

export interface ObserveVote {
  index: string;
  creator: string;
  txhash: string;
  from: string;
  to: string;
  amount: string;
  txtime: string;
  chainId: string;
  used: string;
}

const baseObserveVote: object = {
  index: "",
  creator: "",
  txhash: "",
  from: "",
  to: "",
  amount: "",
  txtime: "",
  chainId: "",
  used: "",
};

export const ObserveVote = {
  encode(message: ObserveVote, writer: Writer = Writer.create()): Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    if (message.creator !== "") {
      writer.uint32(18).string(message.creator);
    }
    if (message.txhash !== "") {
      writer.uint32(26).string(message.txhash);
    }
    if (message.from !== "") {
      writer.uint32(34).string(message.from);
    }
    if (message.to !== "") {
      writer.uint32(42).string(message.to);
    }
    if (message.amount !== "") {
      writer.uint32(50).string(message.amount);
    }
    if (message.txtime !== "") {
      writer.uint32(58).string(message.txtime);
    }
    if (message.chainId !== "") {
      writer.uint32(66).string(message.chainId);
    }
    if (message.used !== "") {
      writer.uint32(74).string(message.used);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): ObserveVote {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseObserveVote } as ObserveVote;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        case 2:
          message.creator = reader.string();
          break;
        case 3:
          message.txhash = reader.string();
          break;
        case 4:
          message.from = reader.string();
          break;
        case 5:
          message.to = reader.string();
          break;
        case 6:
          message.amount = reader.string();
          break;
        case 7:
          message.txtime = reader.string();
          break;
        case 8:
          message.chainId = reader.string();
          break;
        case 9:
          message.used = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ObserveVote {
    const message = { ...baseObserveVote } as ObserveVote;
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.txhash !== undefined && object.txhash !== null) {
      message.txhash = String(object.txhash);
    } else {
      message.txhash = "";
    }
    if (object.from !== undefined && object.from !== null) {
      message.from = String(object.from);
    } else {
      message.from = "";
    }
    if (object.to !== undefined && object.to !== null) {
      message.to = String(object.to);
    } else {
      message.to = "";
    }
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = String(object.amount);
    } else {
      message.amount = "";
    }
    if (object.txtime !== undefined && object.txtime !== null) {
      message.txtime = String(object.txtime);
    } else {
      message.txtime = "";
    }
    if (object.chainId !== undefined && object.chainId !== null) {
      message.chainId = String(object.chainId);
    } else {
      message.chainId = "";
    }
    if (object.used !== undefined && object.used !== null) {
      message.used = String(object.used);
    } else {
      message.used = "";
    }
    return message;
  },

  toJSON(message: ObserveVote): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    message.creator !== undefined && (obj.creator = message.creator);
    message.txhash !== undefined && (obj.txhash = message.txhash);
    message.from !== undefined && (obj.from = message.from);
    message.to !== undefined && (obj.to = message.to);
    message.amount !== undefined && (obj.amount = message.amount);
    message.txtime !== undefined && (obj.txtime = message.txtime);
    message.chainId !== undefined && (obj.chainId = message.chainId);
    message.used !== undefined && (obj.used = message.used);
    return obj;
  },

  fromPartial(object: DeepPartial<ObserveVote>): ObserveVote {
    const message = { ...baseObserveVote } as ObserveVote;
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.txhash !== undefined && object.txhash !== null) {
      message.txhash = object.txhash;
    } else {
      message.txhash = "";
    }
    if (object.from !== undefined && object.from !== null) {
      message.from = object.from;
    } else {
      message.from = "";
    }
    if (object.to !== undefined && object.to !== null) {
      message.to = object.to;
    } else {
      message.to = "";
    }
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = object.amount;
    } else {
      message.amount = "";
    }
    if (object.txtime !== undefined && object.txtime !== null) {
      message.txtime = object.txtime;
    } else {
      message.txtime = "";
    }
    if (object.chainId !== undefined && object.chainId !== null) {
      message.chainId = object.chainId;
    } else {
      message.chainId = "";
    }
    if (object.used !== undefined && object.used !== null) {
      message.used = object.used;
    } else {
      message.used = "";
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
