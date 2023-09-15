import { FUNCTION_FREQUENCY, FUNCTION_GROWTH, FUNCTION_MARGIN, FUNCTION_ROUTE, HOME, SIGIN_ROUTE, TABLE_ROUTE, VIEW_ROUTE } from "./utils/const"
import SigIn from "./pages/SigIn"
import PersonalInfo from "./pages/PersonalInfo"
import Card from "./pages/Card"
import SkuGroup from "./pages/SkuGroup"
import ProductGrid from "./pages/ProductGrid"
import Store from "./pages/Store"
import Transaction from "./pages/Transaction"
import Check from "./pages/Check"
import DateAnalysing from "./pages/DateAnalysing"
import Customers from "./pages/Customers"
import PurchaseHistory from "./pages/PurchaseHistory"
import Periods from "./pages/Periods"
import Groups from "./pages/Groups"
import GrowthAverage from "./pages/GrowthAverage"
import FrequencyVisits from "./pages/FrequencyVisits"
import OfferIncreasingMargin from "./pages/OfferIncreasingMargin"
import Home from "./pages/Home"

export const authRoutes = [
    {
        path: TABLE_ROUTE+"/personal_info",
        Component: PersonalInfo
    },
    {
        path: TABLE_ROUTE+"/cards",
        Component: Card
    },
    {
        path: TABLE_ROUTE+"/sku_group",
        Component: SkuGroup
    },
    {
        path: TABLE_ROUTE+"/product_grid",
        Component: ProductGrid
    },
    {
        path: TABLE_ROUTE+"/stores",
        Component: Store
    },
    {
        path: TABLE_ROUTE+"/transaction",
        Component: Transaction
    },
    {
        path: TABLE_ROUTE+"/checks",
        Component: Check
    },
    {
        path: TABLE_ROUTE+"/date_of_analysing_formation",
        Component: DateAnalysing
    },
    {
        path: VIEW_ROUTE+"/customers",
        Component: Customers
    },
    {
        path: VIEW_ROUTE+"/purchase_history",
        Component: PurchaseHistory
    },
    {
        path: VIEW_ROUTE+"/periods",
        Component: Periods
    },
    {
        path: VIEW_ROUTE+"/groups",
        Component: Groups
    },
    {
        path: FUNCTION_ROUTE+"/"+FUNCTION_GROWTH,
        Component: GrowthAverage
    },
    {
        path: FUNCTION_ROUTE+"/"+FUNCTION_FREQUENCY,
        Component: FrequencyVisits
    },
    {
        path: FUNCTION_ROUTE+"/"+FUNCTION_MARGIN,
        Component: OfferIncreasingMargin
    },
]

export const publicRoutes = [
    {
        path: SIGIN_ROUTE,
        Component: SigIn
    },
    {
        path: HOME,
        Component: Home
    },
]
