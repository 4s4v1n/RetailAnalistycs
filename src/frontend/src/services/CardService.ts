import $api from "../http";
import { AxiosResponse } from "axios";
import { CardResponse } from "../models/response/CardResponse";

const tableName = "cards"

export default class CardService {
    static async get(): Promise<CardResponse[]> {
        const response = $api.get<CardResponse[]>(`/table/${tableName}`)
        return (await response).data;
    }

    static async post(data: CardResponse): Promise<AxiosResponse> {
        data.customer_card_id = Number(data.customer_card_id);
        data.customer_id = Number(data.customer_id);
        const response = $api.post(`/table/${tableName}`, data)
        return response
    }

    static async patch(data: CardResponse): Promise<AxiosResponse> {
        data.customer_id = Number(data.customer_id);
        data.customer_card_id = Number(data.customer_card_id);
        const response = $api.patch(`/table/${tableName}`, data)
        return response
    }

    static async delete(data: string): Promise<AxiosResponse> {
        const response = $api.delete(`/table/${tableName}/${data}`)
        return response
    }

    static async export(): Promise<AxiosResponse> {
        const response = $api.get(`/export/${tableName}`)
        return response
    }

    static async import(data: any): Promise<AxiosResponse> {
        for (let i = 0; i < data.length; i++) {
            data[i].customer_id = Number(data[i].customer_id);
            data[i].customer_card_id = Number(data[i].customer_card_id);
        }
        const response = $api.post(`/import/${tableName}`, data)
        return response
    }
}
