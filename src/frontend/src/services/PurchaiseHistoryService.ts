import $api from "../http";
import { AxiosResponse } from "axios";
import { PurchaiseHistoryResponse } from "../models/response/PurchaiseHistoryResponse";

const tableName = "purchase_history"

export default class PurchaiseHistoryService {
    static async get(): Promise<PurchaiseHistoryResponse[]> {
        const response = $api.get<PurchaiseHistoryResponse[]>(`/view/${tableName}`)
        return (await response).data;
    }

    static async export(): Promise<AxiosResponse> {
        const response = $api.get(`/export/${tableName}`)
        return response
    }
}