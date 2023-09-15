import $api from "../http";
import { AxiosResponse } from "axios";
import { FUNCTION_MARGIN } from "../utils/const";
import { OfferIncreasingMarginRequest } from "../models/request/OfferIncreasingMarginRequest";
import { OfferIncreasingMarginResponse } from "../models/response/OfferIncreasingMarginResponse";

export default class OfferIncreasingMarginService {
    static async get(data: OfferIncreasingMarginRequest): Promise<OfferIncreasingMarginResponse[]> {
        const response = $api.get<OfferIncreasingMarginResponse[]>(`/function/${FUNCTION_MARGIN}`, {
            params: {
                "count_group": Number(data.count_group),
                "max_churn_rate": Number(data.max_churn_rate),
                "max_stability_index": Number(data.max_stability_index),
                "max_index_sku": Number(data.max_index_sku),
                "margin_share": Number(data.margin_share),
            },
            headers: {
                "Accept": "application/json"
            }
        })
        console.log(response)
        return (await response).data;
    }

    static async export(data: OfferIncreasingMarginRequest): Promise<AxiosResponse> {
        const response = $api.get<OfferIncreasingMarginResponse[]>(`/function/${FUNCTION_MARGIN}`, {
            params: {
                "count_group": Number(data.count_group),
                "max_churn_rate": Number(data.max_churn_rate),
                "max_stability_index": Number(data.max_stability_index),
                "max_index_sku": Number(data.max_index_sku),
                "margin_share": Number(data.margin_share),
            },
            headers: {
                "Accept": "text/csv"
            }
        })
        return response;
    }
}
