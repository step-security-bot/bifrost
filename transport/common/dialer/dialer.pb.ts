/* eslint-disable */
import Long from 'long'
import { Backoff } from '../../../util/backoff/backoff.pb.js'
import * as _m0 from 'protobufjs/minimal'

export const protobufPackage = 'dialer'

/** DialerOpts contains options relating to dialing a statically configured peer. */
export interface DialerOpts {
  /** Address is the address of the peer, in the format expected by the transport. */
  address: string
  /**
   * Backoff is the dialing backoff configuration.
   * Can be empty.
   */
  backoff: Backoff | undefined
}

function createBaseDialerOpts(): DialerOpts {
  return { address: '', backoff: undefined }
}

export const DialerOpts = {
  encode(
    message: DialerOpts,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.address !== '') {
      writer.uint32(10).string(message.address)
    }
    if (message.backoff !== undefined) {
      Backoff.encode(message.backoff, writer.uint32(18).fork()).ldelim()
    }
    return writer
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DialerOpts {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input)
    let end = length === undefined ? reader.len : reader.pos + length
    const message = createBaseDialerOpts()
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.address = reader.string()
          break
        case 2:
          message.backoff = Backoff.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  // encodeTransform encodes a source of message objects.
  // Transform<DialerOpts, Uint8Array>
  async *encodeTransform(
    source:
      | AsyncIterable<DialerOpts | DialerOpts[]>
      | Iterable<DialerOpts | DialerOpts[]>
  ): AsyncIterable<Uint8Array> {
    for await (const pkt of source) {
      if (Array.isArray(pkt)) {
        for (const p of pkt) {
          yield* [DialerOpts.encode(p).finish()]
        }
      } else {
        yield* [DialerOpts.encode(pkt).finish()]
      }
    }
  },

  // decodeTransform decodes a source of encoded messages.
  // Transform<Uint8Array, DialerOpts>
  async *decodeTransform(
    source:
      | AsyncIterable<Uint8Array | Uint8Array[]>
      | Iterable<Uint8Array | Uint8Array[]>
  ): AsyncIterable<DialerOpts> {
    for await (const pkt of source) {
      if (Array.isArray(pkt)) {
        for (const p of pkt) {
          yield* [DialerOpts.decode(p)]
        }
      } else {
        yield* [DialerOpts.decode(pkt)]
      }
    }
  },

  fromJSON(object: any): DialerOpts {
    return {
      address: isSet(object.address) ? String(object.address) : '',
      backoff: isSet(object.backoff)
        ? Backoff.fromJSON(object.backoff)
        : undefined,
    }
  },

  toJSON(message: DialerOpts): unknown {
    const obj: any = {}
    message.address !== undefined && (obj.address = message.address)
    message.backoff !== undefined &&
      (obj.backoff = message.backoff
        ? Backoff.toJSON(message.backoff)
        : undefined)
    return obj
  },

  fromPartial<I extends Exact<DeepPartial<DialerOpts>, I>>(
    object: I
  ): DialerOpts {
    const message = createBaseDialerOpts()
    message.address = object.address ?? ''
    message.backoff =
      object.backoff !== undefined && object.backoff !== null
        ? Backoff.fromPartial(object.backoff)
        : undefined
    return message
  },
}

type Builtin =
  | Date
  | Function
  | Uint8Array
  | string
  | number
  | boolean
  | undefined

export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Long
  ? string | number | Long
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends { $case: string }
  ? { [K in keyof Omit<T, '$case'>]?: DeepPartial<T[K]> } & {
      $case: T['$case']
    }
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>

type KeysOfUnion<T> = T extends T ? keyof T : never
export type Exact<P, I extends P> = P extends Builtin
  ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & Record<
        Exclude<keyof I, KeysOfUnion<P>>,
        never
      >

if (_m0.util.Long !== Long) {
  _m0.util.Long = Long as any
  _m0.configure()
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined
}