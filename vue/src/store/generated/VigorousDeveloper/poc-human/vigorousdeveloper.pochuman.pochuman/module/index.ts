// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry, OfflineSigner, EncodeObject, DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgRequestTransaction } from "./types/pochuman/tx";
import { MsgTranfserPoolcoin } from "./types/pochuman/tx";
import { MsgKeysignVote } from "./types/pochuman/tx";
import { MsgApproveTransaction } from "./types/pochuman/tx";
import { MsgObservationVote } from "./types/pochuman/tx";
import { MsgUpdateBalance } from "./types/pochuman/tx";


const types = [
  ["/vigorousdeveloper.pochuman.pochuman.MsgRequestTransaction", MsgRequestTransaction],
  ["/vigorousdeveloper.pochuman.pochuman.MsgTranfserPoolcoin", MsgTranfserPoolcoin],
  ["/vigorousdeveloper.pochuman.pochuman.MsgKeysignVote", MsgKeysignVote],
  ["/vigorousdeveloper.pochuman.pochuman.MsgApproveTransaction", MsgApproveTransaction],
  ["/vigorousdeveloper.pochuman.pochuman.MsgObservationVote", MsgObservationVote],
  ["/vigorousdeveloper.pochuman.pochuman.MsgUpdateBalance", MsgUpdateBalance],
  
];
export const MissingWalletError = new Error("wallet is required");

export const registry = new Registry(<any>types);

const defaultFee = {
  amount: [],
  gas: "200000",
};

interface TxClientOptions {
  addr: string
}

interface SignAndBroadcastOptions {
  fee: StdFee,
  memo?: string
}

const txClient = async (wallet: OfflineSigner, { addr: addr }: TxClientOptions = { addr: "http://localhost:26657" }) => {
  if (!wallet) throw MissingWalletError;
  let client;
  if (addr) {
    client = await SigningStargateClient.connectWithSigner(addr, wallet, { registry });
  }else{
    client = await SigningStargateClient.offline( wallet, { registry });
  }
  const { address } = (await wallet.getAccounts())[0];

  return {
    signAndBroadcast: (msgs: EncodeObject[], { fee, memo }: SignAndBroadcastOptions = {fee: defaultFee, memo: ""}) => client.signAndBroadcast(address, msgs, fee,memo),
    msgRequestTransaction: (data: MsgRequestTransaction): EncodeObject => ({ typeUrl: "/vigorousdeveloper.pochuman.pochuman.MsgRequestTransaction", value: MsgRequestTransaction.fromPartial( data ) }),
    msgTranfserPoolcoin: (data: MsgTranfserPoolcoin): EncodeObject => ({ typeUrl: "/vigorousdeveloper.pochuman.pochuman.MsgTranfserPoolcoin", value: MsgTranfserPoolcoin.fromPartial( data ) }),
    msgKeysignVote: (data: MsgKeysignVote): EncodeObject => ({ typeUrl: "/vigorousdeveloper.pochuman.pochuman.MsgKeysignVote", value: MsgKeysignVote.fromPartial( data ) }),
    msgApproveTransaction: (data: MsgApproveTransaction): EncodeObject => ({ typeUrl: "/vigorousdeveloper.pochuman.pochuman.MsgApproveTransaction", value: MsgApproveTransaction.fromPartial( data ) }),
    msgObservationVote: (data: MsgObservationVote): EncodeObject => ({ typeUrl: "/vigorousdeveloper.pochuman.pochuman.MsgObservationVote", value: MsgObservationVote.fromPartial( data ) }),
    msgUpdateBalance: (data: MsgUpdateBalance): EncodeObject => ({ typeUrl: "/vigorousdeveloper.pochuman.pochuman.MsgUpdateBalance", value: MsgUpdateBalance.fromPartial( data ) }),
    
  };
};

interface QueryClientOptions {
  addr: string
}

const queryClient = async ({ addr: addr }: QueryClientOptions = { addr: "http://localhost:1317" }) => {
  return new Api({ baseUrl: addr });
};

export {
  txClient,
  queryClient,
};
