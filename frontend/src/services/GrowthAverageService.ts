import $api from "../http";
import { AxiosResponse } from "axios";
import { GrowthAverageRequest } from "../models/request/GrowthAverageRequest";
import { GrowthAverageResponse } from "../models/response/GrowthAverageResponse";

const tableName = "growth_of_average_check"

export default class GrowthAverageService {
    static async get(data: GrowthAverageRequest): Promise<GrowthAverageResponse[]> {
        const response = $api.get<GrowthAverageResponse[]>(`/function/${tableName}`, {
            params: {
                "method": data.method,
                "first": data.first,
                "last": data.last,
                "number": Number(data.number),
                "coefficient": Number(data.coefficient),
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

    static async export(data: GrowthAverageRequest): Promise<AxiosResponse> {
        const response = $api.get<GrowthAverageResponse[]>(`/function/${tableName}`, {
            params: {
                "method": data.method,
                "first": data.first,
                "last": data.last,
                "number": Number(data.number),
                "coefficient": Number(data.coefficient),
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