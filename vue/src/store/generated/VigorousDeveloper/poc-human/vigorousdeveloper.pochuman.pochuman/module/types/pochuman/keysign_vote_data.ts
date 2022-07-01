/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "vigorousdeveloper.pochuman.pochuman";

export interface KeysignVoteData {
  index: string;
  txHash: string;
  pubKey: string;
  voter: string;
  txTime: string;
}

const baseKeysignVoteData: object = {
  index: "",
  txHash: "",
  pubKey: "",
  voter: "",
  txTime: "",
};

export const KeysignVoteData = {
  encode(message: KeysignVoteData, writer: Writer = Writer.create()): Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    if (message.txHash !== "") {
      writer.uint32(18).string(message.txHash);
    }
    if (message.pubKey !== "") {
      writer.uint32(26).string(message.pubKey);
    }
    if (message.voter !== "") {
      writer.uint32(34).string(message.voter);
    }
    if (message.txTime !== "") {
      writer.uint32(42).string(message.txTime);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): KeysignVoteData {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseKeysignVoteData } as KeysignVoteData;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        case 2:
          message.txHash = reader.string();
          break;
        case 3:
          message.pubKey = reader.string();
          break;
        case 4:
          message.voter = reader.string();
          break;
        case 5:
          message.txTime = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): KeysignVoteData {
    const message = { ...baseKeysignVoteData } as KeysignVoteData;
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    if (object.txHash !== undefined && object.txHash !== null) {
      message.txHash = String(object.txHash);
    } else {
      message.txHash = "";
    }
    if (object.pubKey !== undefined && object.pubKey !== null) {
      message.pubKey = String(object.pubKey);
    } else {
      message.pubKey = "";
    }
    if (object.voter !== undefined && object.voter !== null) {
      message.voter = String(object.voter);
    } else {
      message.voter = "";
    }
    if (object.txTime !== undefined && object.txTime !== null) {
      message.txTime = String(object.txTime);
    } else {
      message.txTime = "";
    }
    return message;
  },

  toJSON(message: KeysignVoteData): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    message.txHash !== undefined && (obj.txHash = message.txHash);
    message.pubKey !== undefined && (obj.pubKey = message.pubKey);
    message.voter !== undefined && (obj.voter = message.voter);
    message.txTime !== undefined && (obj.txTime = message.txTime);
    return obj;
  },

  fromPartial(object: DeepPartial<KeysignVoteData>): KeysignVoteData {
    const message = { ...baseKeysignVoteData } as KeysignVoteData;
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    if (object.txHash !== undefined && object.txHash !== null) {
      message.txHash = object.txHash;
    } else {
      message.txHash = "";
    }
    if (object.pubKey !== undefined && object.pubKey !== null) {
      message.pubKey = object.pubKey;
    } else {
      message.pubKey = "";
    }
    if (object.voter !== undefined && object.voter !== null) {
      message.voter = object.voter;
    } else {
      message.voter = "";
    }
    if (object.txTime !== undefined && object.txTime !== null) {
      message.txTime = object.txTime;
    } else {
      message.txTime = "";
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
