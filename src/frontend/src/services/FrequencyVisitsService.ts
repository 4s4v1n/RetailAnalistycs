import $api from "../http";
import { AxiosResponse } from "axios";
import { FUNCTION_FREQUENCY } from "../utils/const";
import { FrequencyVisitsRequest } from "../models/request/FrequencyVisitsRequest";
import { FrequencyVisitsResponse } from "../models/response/FrequencyVisitsResponse";

export default class FrequencyVisitsService {
    static async get(data: FrequencyVisitsRequest): Promise<FrequencyVisitsResponse[]> {
        const response = $api.get<FrequencyVisitsResponse[]>(`/function/${FUNCTION_FREQUENCY}`, {
            params: {
                "first": data.first,
                "last": data.last,
                "value_transaction": Number(data.value_transaction),
                "max_churn_rate": Number(data.max_churn_rate),
                "max_discount_share": Number(data.max_discount_share),
                "margin_share": Number(data.margin_share),
            },
            headers: {
                "Accept": "application/json"
            }
        })
        return (await response).data;
    }

    static async export(data: FrequencyVisitsRequest): Promise<AxiosResponse> {
        const response = $api.get<FrequencyVisitsResponse[]>(`/function/${FUNCTION_FREQUENCY}`, {
            params: {
                "first": data.first,
                "last": data.last,
                "value_transaction": Number(data.value_transaction),
                "max_churn_rate": Number(data.max_churn_rate),
                "max_discount_share": Number(data.max_discount_share),
                "margin_share": Number(data.margin_share),
            },
            headers: {
                "Accept": "text/csv"
            }
        })
        return response;
    }
}