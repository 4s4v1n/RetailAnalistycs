import { IData } from "../models/IData";
import { FUNCTION_FREQUENCY, FUNCTION_GROWTH, FUNCTION_MARGIN, FUNCTION_ROUTE, TABLE_ROUTE, VIEW_ROUTE } from "../utils/const";


export const DataLabel: IData[] = [
    {
        name: "Personal Information",
        link: TABLE_ROUTE+"/personal_info"
    },
    {
        name: "Cards",
        link: TABLE_ROUTE+"/cards"
    },
    {
        name: "Sku Group",
        link: TABLE_ROUTE+"/sku_group"
    },
    {
        name: "Product Grid",
        link: TABLE_ROUTE+"/product_grid"
    },
    {
        name: "Stores",
        link: TABLE_ROUTE+"/stores",
    },
    {
        name: "Transactions",
        link: TABLE_ROUTE+"/transaction",
    },
    {
        name: "Checks",
        link: TABLE_ROUTE+"/checks",
    },
    {
        name: "Date Analysing",
        link: TABLE_ROUTE+"/date_of_analysing_formation"
    }
];

export const ViewLabel: IData[] = [
    {
        name: "Customers",
        link: VIEW_ROUTE+"/customers"
    },
    {
        name: "Purchase History",
        link: VIEW_ROUTE+"/purchase_history"
    },
    {
        name: "Periods",
        link: VIEW_ROUTE+"/periods"
    },
    {
        name: "Groups",
        link: VIEW_ROUTE+"/groups"
    },
];

export const FunctionLabel: IData[] = [
    {
        name: "Average check",
        link: FUNCTION_ROUTE+"/"+FUNCTION_GROWTH
    },
    {
        name: "Frequency visits",
        link: FUNCTION_ROUTE+"/"+FUNCTION_FREQUENCY
    },
    {
        name: "Increasing margin",
        link: FUNCTION_ROUTE+"/"+FUNCTION_MARGIN
    },
];

