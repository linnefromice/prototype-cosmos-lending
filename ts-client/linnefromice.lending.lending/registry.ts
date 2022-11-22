import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgWithdraw } from "./types/lending/lending/tx";
import { MsgAddPool } from "./types/lending/lending/tx";
import { MsgDeposit } from "./types/lending/lending/tx";
import { MsgBorrow } from "./types/lending/lending/tx";
import { MsgRepay } from "./types/lending/lending/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/linnefromice.lending.lending.MsgWithdraw", MsgWithdraw],
    ["/linnefromice.lending.lending.MsgAddPool", MsgAddPool],
    ["/linnefromice.lending.lending.MsgDeposit", MsgDeposit],
    ["/linnefromice.lending.lending.MsgBorrow", MsgBorrow],
    ["/linnefromice.lending.lending.MsgRepay", MsgRepay],
    
];

export { msgTypes }