// Generated by Ignite ignite.com/cli

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient, DeliverTxResponse } from "@cosmjs/stargate";
import { EncodeObject, GeneratedType, OfflineSigner, Registry } from "@cosmjs/proto-signing";
import { msgTypes } from './registry';
import { IgniteClient } from "../client"
import { MissingWalletError } from "../helpers"
import { Api } from "./rest";
import { MsgSubmitProposal } from "./types/cosmos/group/v1/tx";
import { MsgCreateGroupWithPolicy } from "./types/cosmos/group/v1/tx";
import { MsgWithdrawProposal } from "./types/cosmos/group/v1/tx";
import { MsgCreateGroup } from "./types/cosmos/group/v1/tx";
import { MsgUpdateGroupMembers } from "./types/cosmos/group/v1/tx";
import { MsgUpdateGroupPolicyMetadata } from "./types/cosmos/group/v1/tx";
import { MsgUpdateGroupMetadata } from "./types/cosmos/group/v1/tx";
import { MsgUpdateGroupPolicyAdmin } from "./types/cosmos/group/v1/tx";
import { MsgUpdateGroupAdmin } from "./types/cosmos/group/v1/tx";
import { MsgCreateGroupPolicy } from "./types/cosmos/group/v1/tx";
import { MsgExec } from "./types/cosmos/group/v1/tx";
import { MsgUpdateGroupPolicyDecisionPolicy } from "./types/cosmos/group/v1/tx";
import { MsgLeaveGroup } from "./types/cosmos/group/v1/tx";
import { MsgVote } from "./types/cosmos/group/v1/tx";


export { MsgSubmitProposal, MsgCreateGroupWithPolicy, MsgWithdrawProposal, MsgCreateGroup, MsgUpdateGroupMembers, MsgUpdateGroupPolicyMetadata, MsgUpdateGroupMetadata, MsgUpdateGroupPolicyAdmin, MsgUpdateGroupAdmin, MsgCreateGroupPolicy, MsgExec, MsgUpdateGroupPolicyDecisionPolicy, MsgLeaveGroup, MsgVote };

type sendMsgSubmitProposalParams = {
  value: MsgSubmitProposal,
  fee?: StdFee,
  memo?: string
};

type sendMsgCreateGroupWithPolicyParams = {
  value: MsgCreateGroupWithPolicy,
  fee?: StdFee,
  memo?: string
};

type sendMsgWithdrawProposalParams = {
  value: MsgWithdrawProposal,
  fee?: StdFee,
  memo?: string
};

type sendMsgCreateGroupParams = {
  value: MsgCreateGroup,
  fee?: StdFee,
  memo?: string
};

type sendMsgUpdateGroupMembersParams = {
  value: MsgUpdateGroupMembers,
  fee?: StdFee,
  memo?: string
};

type sendMsgUpdateGroupPolicyMetadataParams = {
  value: MsgUpdateGroupPolicyMetadata,
  fee?: StdFee,
  memo?: string
};

type sendMsgUpdateGroupMetadataParams = {
  value: MsgUpdateGroupMetadata,
  fee?: StdFee,
  memo?: string
};

type sendMsgUpdateGroupPolicyAdminParams = {
  value: MsgUpdateGroupPolicyAdmin,
  fee?: StdFee,
  memo?: string
};

type sendMsgUpdateGroupAdminParams = {
  value: MsgUpdateGroupAdmin,
  fee?: StdFee,
  memo?: string
};

type sendMsgCreateGroupPolicyParams = {
  value: MsgCreateGroupPolicy,
  fee?: StdFee,
  memo?: string
};

type sendMsgExecParams = {
  value: MsgExec,
  fee?: StdFee,
  memo?: string
};

type sendMsgUpdateGroupPolicyDecisionPolicyParams = {
  value: MsgUpdateGroupPolicyDecisionPolicy,
  fee?: StdFee,
  memo?: string
};

type sendMsgLeaveGroupParams = {
  value: MsgLeaveGroup,
  fee?: StdFee,
  memo?: string
};

type sendMsgVoteParams = {
  value: MsgVote,
  fee?: StdFee,
  memo?: string
};


type msgSubmitProposalParams = {
  value: MsgSubmitProposal,
};

type msgCreateGroupWithPolicyParams = {
  value: MsgCreateGroupWithPolicy,
};

type msgWithdrawProposalParams = {
  value: MsgWithdrawProposal,
};

type msgCreateGroupParams = {
  value: MsgCreateGroup,
};

type msgUpdateGroupMembersParams = {
  value: MsgUpdateGroupMembers,
};

type msgUpdateGroupPolicyMetadataParams = {
  value: MsgUpdateGroupPolicyMetadata,
};

type msgUpdateGroupMetadataParams = {
  value: MsgUpdateGroupMetadata,
};

type msgUpdateGroupPolicyAdminParams = {
  value: MsgUpdateGroupPolicyAdmin,
};

type msgUpdateGroupAdminParams = {
  value: MsgUpdateGroupAdmin,
};

type msgCreateGroupPolicyParams = {
  value: MsgCreateGroupPolicy,
};

type msgExecParams = {
  value: MsgExec,
};

type msgUpdateGroupPolicyDecisionPolicyParams = {
  value: MsgUpdateGroupPolicyDecisionPolicy,
};

type msgLeaveGroupParams = {
  value: MsgLeaveGroup,
};

type msgVoteParams = {
  value: MsgVote,
};


export const registry = new Registry(msgTypes);

const defaultFee = {
  amount: [],
  gas: "200000",
};

interface TxClientOptions {
  addr: string
	prefix: string
	signer?: OfflineSigner
}

export const txClient = ({ signer, prefix, addr }: TxClientOptions = { addr: "http://localhost:26657", prefix: "cosmos" }) => {

  return {
		
		async sendMsgSubmitProposal({ value, fee, memo }: sendMsgSubmitProposalParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgSubmitProposal: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgSubmitProposal({ value: MsgSubmitProposal.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgSubmitProposal: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgCreateGroupWithPolicy({ value, fee, memo }: sendMsgCreateGroupWithPolicyParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgCreateGroupWithPolicy: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgCreateGroupWithPolicy({ value: MsgCreateGroupWithPolicy.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgCreateGroupWithPolicy: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgWithdrawProposal({ value, fee, memo }: sendMsgWithdrawProposalParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgWithdrawProposal: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgWithdrawProposal({ value: MsgWithdrawProposal.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgWithdrawProposal: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgCreateGroup({ value, fee, memo }: sendMsgCreateGroupParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgCreateGroup: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgCreateGroup({ value: MsgCreateGroup.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgCreateGroup: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgUpdateGroupMembers({ value, fee, memo }: sendMsgUpdateGroupMembersParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgUpdateGroupMembers: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgUpdateGroupMembers({ value: MsgUpdateGroupMembers.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgUpdateGroupMembers: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgUpdateGroupPolicyMetadata({ value, fee, memo }: sendMsgUpdateGroupPolicyMetadataParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgUpdateGroupPolicyMetadata: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgUpdateGroupPolicyMetadata({ value: MsgUpdateGroupPolicyMetadata.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgUpdateGroupPolicyMetadata: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgUpdateGroupMetadata({ value, fee, memo }: sendMsgUpdateGroupMetadataParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgUpdateGroupMetadata: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgUpdateGroupMetadata({ value: MsgUpdateGroupMetadata.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgUpdateGroupMetadata: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgUpdateGroupPolicyAdmin({ value, fee, memo }: sendMsgUpdateGroupPolicyAdminParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgUpdateGroupPolicyAdmin: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgUpdateGroupPolicyAdmin({ value: MsgUpdateGroupPolicyAdmin.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgUpdateGroupPolicyAdmin: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgUpdateGroupAdmin({ value, fee, memo }: sendMsgUpdateGroupAdminParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgUpdateGroupAdmin: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgUpdateGroupAdmin({ value: MsgUpdateGroupAdmin.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgUpdateGroupAdmin: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgCreateGroupPolicy({ value, fee, memo }: sendMsgCreateGroupPolicyParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgCreateGroupPolicy: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgCreateGroupPolicy({ value: MsgCreateGroupPolicy.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgCreateGroupPolicy: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgExec({ value, fee, memo }: sendMsgExecParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgExec: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgExec({ value: MsgExec.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgExec: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgUpdateGroupPolicyDecisionPolicy({ value, fee, memo }: sendMsgUpdateGroupPolicyDecisionPolicyParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgUpdateGroupPolicyDecisionPolicy: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgUpdateGroupPolicyDecisionPolicy({ value: MsgUpdateGroupPolicyDecisionPolicy.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgUpdateGroupPolicyDecisionPolicy: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgLeaveGroup({ value, fee, memo }: sendMsgLeaveGroupParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgLeaveGroup: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgLeaveGroup({ value: MsgLeaveGroup.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgLeaveGroup: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgVote({ value, fee, memo }: sendMsgVoteParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgVote: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgVote({ value: MsgVote.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgVote: Could not broadcast Tx: '+ e.message)
			}
		},
		
		
		msgSubmitProposal({ value }: msgSubmitProposalParams): EncodeObject {
			try {
				return { typeUrl: "/cosmos.group.v1.MsgSubmitProposal", value: MsgSubmitProposal.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgSubmitProposal: Could not create message: ' + e.message)
			}
		},
		
		msgCreateGroupWithPolicy({ value }: msgCreateGroupWithPolicyParams): EncodeObject {
			try {
				return { typeUrl: "/cosmos.group.v1.MsgCreateGroupWithPolicy", value: MsgCreateGroupWithPolicy.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgCreateGroupWithPolicy: Could not create message: ' + e.message)
			}
		},
		
		msgWithdrawProposal({ value }: msgWithdrawProposalParams): EncodeObject {
			try {
				return { typeUrl: "/cosmos.group.v1.MsgWithdrawProposal", value: MsgWithdrawProposal.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgWithdrawProposal: Could not create message: ' + e.message)
			}
		},
		
		msgCreateGroup({ value }: msgCreateGroupParams): EncodeObject {
			try {
				return { typeUrl: "/cosmos.group.v1.MsgCreateGroup", value: MsgCreateGroup.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgCreateGroup: Could not create message: ' + e.message)
			}
		},
		
		msgUpdateGroupMembers({ value }: msgUpdateGroupMembersParams): EncodeObject {
			try {
				return { typeUrl: "/cosmos.group.v1.MsgUpdateGroupMembers", value: MsgUpdateGroupMembers.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgUpdateGroupMembers: Could not create message: ' + e.message)
			}
		},
		
		msgUpdateGroupPolicyMetadata({ value }: msgUpdateGroupPolicyMetadataParams): EncodeObject {
			try {
				return { typeUrl: "/cosmos.group.v1.MsgUpdateGroupPolicyMetadata", value: MsgUpdateGroupPolicyMetadata.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgUpdateGroupPolicyMetadata: Could not create message: ' + e.message)
			}
		},
		
		msgUpdateGroupMetadata({ value }: msgUpdateGroupMetadataParams): EncodeObject {
			try {
				return { typeUrl: "/cosmos.group.v1.MsgUpdateGroupMetadata", value: MsgUpdateGroupMetadata.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgUpdateGroupMetadata: Could not create message: ' + e.message)
			}
		},
		
		msgUpdateGroupPolicyAdmin({ value }: msgUpdateGroupPolicyAdminParams): EncodeObject {
			try {
				return { typeUrl: "/cosmos.group.v1.MsgUpdateGroupPolicyAdmin", value: MsgUpdateGroupPolicyAdmin.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgUpdateGroupPolicyAdmin: Could not create message: ' + e.message)
			}
		},
		
		msgUpdateGroupAdmin({ value }: msgUpdateGroupAdminParams): EncodeObject {
			try {
				return { typeUrl: "/cosmos.group.v1.MsgUpdateGroupAdmin", value: MsgUpdateGroupAdmin.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgUpdateGroupAdmin: Could not create message: ' + e.message)
			}
		},
		
		msgCreateGroupPolicy({ value }: msgCreateGroupPolicyParams): EncodeObject {
			try {
				return { typeUrl: "/cosmos.group.v1.MsgCreateGroupPolicy", value: MsgCreateGroupPolicy.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgCreateGroupPolicy: Could not create message: ' + e.message)
			}
		},
		
		msgExec({ value }: msgExecParams): EncodeObject {
			try {
				return { typeUrl: "/cosmos.group.v1.MsgExec", value: MsgExec.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgExec: Could not create message: ' + e.message)
			}
		},
		
		msgUpdateGroupPolicyDecisionPolicy({ value }: msgUpdateGroupPolicyDecisionPolicyParams): EncodeObject {
			try {
				return { typeUrl: "/cosmos.group.v1.MsgUpdateGroupPolicyDecisionPolicy", value: MsgUpdateGroupPolicyDecisionPolicy.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgUpdateGroupPolicyDecisionPolicy: Could not create message: ' + e.message)
			}
		},
		
		msgLeaveGroup({ value }: msgLeaveGroupParams): EncodeObject {
			try {
				return { typeUrl: "/cosmos.group.v1.MsgLeaveGroup", value: MsgLeaveGroup.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgLeaveGroup: Could not create message: ' + e.message)
			}
		},
		
		msgVote({ value }: msgVoteParams): EncodeObject {
			try {
				return { typeUrl: "/cosmos.group.v1.MsgVote", value: MsgVote.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgVote: Could not create message: ' + e.message)
			}
		},
		
	}
};

interface QueryClientOptions {
  addr: string
}

export const queryClient = ({ addr: addr }: QueryClientOptions = { addr: "http://localhost:1317" }) => {
  return new Api({ baseURL: addr });
};

class SDKModule {
	public query: ReturnType<typeof queryClient>;
	public tx: ReturnType<typeof txClient>;
	
	public registry: Array<[string, GeneratedType]> = [];

	constructor(client: IgniteClient) {		
	
		this.query = queryClient({ addr: client.env.apiURL });		
		this.updateTX(client);
		client.on('signer-changed',(signer) => {			
		 this.updateTX(client);
		})
	}
	updateTX(client: IgniteClient) {
    const methods = txClient({
        signer: client.signer,
        addr: client.env.rpcURL,
        prefix: client.env.prefix ?? "cosmos",
    })
	
    this.tx = methods;
    for (let m in methods) {
        this.tx[m] = methods[m].bind(this.tx);
    }
	}
};

const Module = (test: IgniteClient) => {
	return {
		module: {
			CosmosGroupV1: new SDKModule(test)
		},
		registry: msgTypes
  }
}
export default Module;