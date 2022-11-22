import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgBorrow } from "./types/lending/lending/tx";
import { MsgWithdraw } from "./types/lending/lending/tx";
import { MsgAddPool } from "./types/lending/lending/tx";
import { MsgDeposit } from "./types/lending/lending/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/linnefromice.lending.lending.MsgBorrow", MsgBorrow],
    ["/linnefromice.lending.lending.MsgWithdraw", MsgWithdraw],
    ["/linnefromice.lending.lending.MsgAddPool", MsgAddPool],
    ["/linnefromice.lending.lending.MsgDeposit", MsgDeposit],
    
];

export { msgTypes }