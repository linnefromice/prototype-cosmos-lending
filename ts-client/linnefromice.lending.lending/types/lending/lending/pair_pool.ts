/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";
import { Coin } from "../../cosmos/base/v1beta1/coin";

export const protobufPackage = "linnefromice.lending.lending";

export interface PairPool {
  id: number;
  address: string;
  poolId: number;
  assetLiquidity: Coin | undefined;
  assetLpCoinDenom: string;
  assetTotalNormalDeposited: number;
  assetTotalConlyDeposited: number;
  assetTotalBorrowed: number;
  shadowLiquidity: Coin | undefined;
  shadowLpCoinDenom: string;
  shadowTotalNormalDeposited: number;
  shadowTotalConlyDeposited: number;
  shadowTotalBorrowed: number;
  lastUpdated: number;
}

function createBasePairPool(): PairPool {
  return {
    id: 0,
    address: "",
    poolId: 0,
    assetLiquidity: undefined,
    assetLpCoinDenom: "",
    assetTotalNormalDeposited: 0,
    assetTotalConlyDeposited: 0,
    assetTotalBorrowed: 0,
    shadowLiquidity: undefined,
    shadowLpCoinDenom: "",
    shadowTotalNormalDeposited: 0,
    shadowTotalConlyDeposited: 0,
    shadowTotalBorrowed: 0,
    lastUpdated: 0,
  };
}

export const PairPool = {
  encode(message: PairPool, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    if (message.address !== "") {
      writer.uint32(18).string(message.address);
    }
    if (message.poolId !== 0) {
      writer.uint32(24).uint64(message.poolId);
    }
    if (message.assetLiquidity !== undefined) {
      Coin.encode(message.assetLiquidity, writer.uint32(34).fork()).ldelim();
    }
    if (message.assetLpCoinDenom !== "") {
      writer.uint32(42).string(message.assetLpCoinDenom);
    }
    if (message.assetTotalNormalDeposited !== 0) {
      writer.uint32(48).uint64(message.assetTotalNormalDeposited);
    }
    if (message.assetTotalConlyDeposited !== 0) {
      writer.uint32(56).uint64(message.assetTotalConlyDeposited);
    }
    if (message.assetTotalBorrowed !== 0) {
      writer.uint32(64).uint64(message.assetTotalBorrowed);
    }
    if (message.shadowLiquidity !== undefined) {
      Coin.encode(message.shadowLiquidity, writer.uint32(74).fork()).ldelim();
    }
    if (message.shadowLpCoinDenom !== "") {
      writer.uint32(82).string(message.shadowLpCoinDenom);
    }
    if (message.shadowTotalNormalDeposited !== 0) {
      writer.uint32(88).uint64(message.shadowTotalNormalDeposited);
    }
    if (message.shadowTotalConlyDeposited !== 0) {
      writer.uint32(96).uint64(message.shadowTotalConlyDeposited);
    }
    if (message.shadowTotalBorrowed !== 0) {
      writer.uint32(104).uint64(message.shadowTotalBorrowed);
    }
    if (message.lastUpdated !== 0) {
      writer.uint32(112).uint64(message.lastUpdated);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): PairPool {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBasePairPool();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.address = reader.string();
          break;
        case 3:
          message.poolId = longToNumber(reader.uint64() as Long);
          break;
        case 4:
          message.assetLiquidity = Coin.decode(reader, reader.uint32());
          break;
        case 5:
          message.assetLpCoinDenom = reader.string();
          break;
        case 6:
          message.assetTotalNormalDeposited = longToNumber(reader.uint64() as Long);
          break;
        case 7:
          message.assetTotalConlyDeposited = longToNumber(reader.uint64() as Long);
          break;
        case 8:
          message.assetTotalBorrowed = longToNumber(reader.uint64() as Long);
          break;
        case 9:
          message.shadowLiquidity = Coin.decode(reader, reader.uint32());
          break;
        case 10:
          message.shadowLpCoinDenom = reader.string();
          break;
        case 11:
          message.shadowTotalNormalDeposited = longToNumber(reader.uint64() as Long);
          break;
        case 12:
          message.shadowTotalConlyDeposited = longToNumber(reader.uint64() as Long);
          break;
        case 13:
          message.shadowTotalBorrowed = longToNumber(reader.uint64() as Long);
          break;
        case 14:
          message.lastUpdated = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): PairPool {
    return {
      id: isSet(object.id) ? Number(object.id) : 0,
      address: isSet(object.address) ? String(object.address) : "",
      poolId: isSet(object.poolId) ? Number(object.poolId) : 0,
      assetLiquidity: isSet(object.assetLiquidity) ? Coin.fromJSON(object.assetLiquidity) : undefined,
      assetLpCoinDenom: isSet(object.assetLpCoinDenom) ? String(object.assetLpCoinDenom) : "",
      assetTotalNormalDeposited: isSet(object.assetTotalNormalDeposited) ? Number(object.assetTotalNormalDeposited) : 0,
      assetTotalConlyDeposited: isSet(object.assetTotalConlyDeposited) ? Number(object.assetTotalConlyDeposited) : 0,
      assetTotalBorrowed: isSet(object.assetTotalBorrowed) ? Number(object.assetTotalBorrowed) : 0,
      shadowLiquidity: isSet(object.shadowLiquidity) ? Coin.fromJSON(object.shadowLiquidity) : undefined,
      shadowLpCoinDenom: isSet(object.shadowLpCoinDenom) ? String(object.shadowLpCoinDenom) : "",
      shadowTotalNormalDeposited: isSet(object.shadowTotalNormalDeposited)
        ? Number(object.shadowTotalNormalDeposited)
        : 0,
      shadowTotalConlyDeposited: isSet(object.shadowTotalConlyDeposited) ? Number(object.shadowTotalConlyDeposited) : 0,
      shadowTotalBorrowed: isSet(object.shadowTotalBorrowed) ? Number(object.shadowTotalBorrowed) : 0,
      lastUpdated: isSet(object.lastUpdated) ? Number(object.lastUpdated) : 0,
    };
  },

  toJSON(message: PairPool): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = Math.round(message.id));
    message.address !== undefined && (obj.address = message.address);
    message.poolId !== undefined && (obj.poolId = Math.round(message.poolId));
    message.assetLiquidity !== undefined
      && (obj.assetLiquidity = message.assetLiquidity ? Coin.toJSON(message.assetLiquidity) : undefined);
    message.assetLpCoinDenom !== undefined && (obj.assetLpCoinDenom = message.assetLpCoinDenom);
    message.assetTotalNormalDeposited !== undefined
      && (obj.assetTotalNormalDeposited = Math.round(message.assetTotalNormalDeposited));
    message.assetTotalConlyDeposited !== undefined
      && (obj.assetTotalConlyDeposited = Math.round(message.assetTotalConlyDeposited));
    message.assetTotalBorrowed !== undefined && (obj.assetTotalBorrowed = Math.round(message.assetTotalBorrowed));
    message.shadowLiquidity !== undefined
      && (obj.shadowLiquidity = message.shadowLiquidity ? Coin.toJSON(message.shadowLiquidity) : undefined);
    message.shadowLpCoinDenom !== undefined && (obj.shadowLpCoinDenom = message.shadowLpCoinDenom);
    message.shadowTotalNormalDeposited !== undefined
      && (obj.shadowTotalNormalDeposited = Math.round(message.shadowTotalNormalDeposited));
    message.shadowTotalConlyDeposited !== undefined
      && (obj.shadowTotalConlyDeposited = Math.round(message.shadowTotalConlyDeposited));
    message.shadowTotalBorrowed !== undefined && (obj.shadowTotalBorrowed = Math.round(message.shadowTotalBorrowed));
    message.lastUpdated !== undefined && (obj.lastUpdated = Math.round(message.lastUpdated));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<PairPool>, I>>(object: I): PairPool {
    const message = createBasePairPool();
    message.id = object.id ?? 0;
    message.address = object.address ?? "";
    message.poolId = object.poolId ?? 0;
    message.assetLiquidity = (object.assetLiquidity !== undefined && object.assetLiquidity !== null)
      ? Coin.fromPartial(object.assetLiquidity)
      : undefined;
    message.assetLpCoinDenom = object.assetLpCoinDenom ?? "";
    message.assetTotalNormalDeposited = object.assetTotalNormalDeposited ?? 0;
    message.assetTotalConlyDeposited = object.assetTotalConlyDeposited ?? 0;
    message.assetTotalBorrowed = object.assetTotalBorrowed ?? 0;
    message.shadowLiquidity = (object.shadowLiquidity !== undefined && object.shadowLiquidity !== null)
      ? Coin.fromPartial(object.shadowLiquidity)
      : undefined;
    message.shadowLpCoinDenom = object.shadowLpCoinDenom ?? "";
    message.shadowTotalNormalDeposited = object.shadowTotalNormalDeposited ?? 0;
    message.shadowTotalConlyDeposited = object.shadowTotalConlyDeposited ?? 0;
    message.shadowTotalBorrowed = object.shadowTotalBorrowed ?? 0;
    message.lastUpdated = object.lastUpdated ?? 0;
    return message;
  },
};

declare var self: any | undefined;
declare var window: any | undefined;
declare var global: any | undefined;
var globalThis: any = (() => {
  if (typeof globalThis !== "undefined") {
    return globalThis;
  }
  if (typeof self !== "undefined") {
    return self;
  }
  if (typeof window !== "undefined") {
    return window;
  }
  if (typeof global !== "undefined") {
    return global;
  }
  throw "Unable to locate global object";
})();

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (_m0.util.Long !== Long) {
  _m0.util.Long = Long as any;
  _m0.configure();
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
