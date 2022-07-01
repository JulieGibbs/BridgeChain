/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "vigorousdeveloper.pochuman.pochuman";

export interface FeeBalance {
  index: string;
  chainName: string;
  balance: string;
  decimal: string;
}

const baseFeeBalance: object = {
  index: "",
  chainName: "",
  balance: "",
  decimal: "",
};

export const FeeBalance = {
  encode(message: FeeBalance, writer: Writer = Writer.create()): Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    if (message.chainName !== "") {
      writer.uint32(18).string(message.chainName);
    }
    if (message.balance !== "") {
      writer.uint32(26).string(message.balance);
    }
    if (message.decimal !== "") {
      writer.uint32(34).string(message.decimal);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): FeeBalance {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseFeeBalance } as FeeBalance;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        case 2:
          message.chainName = reader.string();
          break;
        case 3:
          message.balance = reader.string();
          break;
        case 4:
          message.decimal = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): FeeBalance {
    const message = { ...baseFeeBalance } as FeeBalance;
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    if (object.chainName !== undefined && object.chainName !== null) {
      message.chainName = String(object.chainName);
    } else {
      message.chainName = "";
    }
    if (object.balance !== undefined && object.balance !== null) {
      message.balance = String(object.balance);
    } else {
      message.balance = "";
    }
    if (object.decimal !== undefined && object.decimal !== null) {
      message.decimal = String(object.decimal);
    } else {
      message.decimal = "";
    }
    return message;
  },

  toJSON(message: FeeBalance): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    message.chainName !== undefined && (obj.chainName = message.chainName);
    message.balance !== undefined && (obj.balance = message.balance);
    message.decimal !== undefined && (obj.decimal = message.decimal);
    return obj;
  },

  fromPartial(object: DeepPartial<FeeBalance>): FeeBalance {
    const message = { ...baseFeeBalance } as FeeBalance;
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    if (object.chainName !== undefined && object.chainName !== null) {
      message.chainName = object.chainName;
    } else {
      message.chainName = "";
    }
    if (object.balance !== undefined && object.balance !== null) {
      message.balance = object.balance;
    } else {
      message.balance = "";
    }
    if (object.decimal !== undefined && object.decimal !== null) {
      message.decimal = object.decimal;
    } else {
      message.decimal = "";
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
