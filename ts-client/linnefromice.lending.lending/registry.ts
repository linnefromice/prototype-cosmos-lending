import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgDeposit } from "./types/lending/lending/tx";
import { MsgAddPool } from "./types/lending/lending/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/linnefromice.lending.lending.MsgDeposit", MsgDeposit],
    ["/linnefromice.lending.lending.MsgAddPool", MsgAddPool],
    
];

export { msgTypes }