import { makeAutoObservable } from "mobx";
import AuthService from "../services/AuthService";
import axios from "axios";
import { AuthResponse } from "../models/response/AuthResponse";
import { API_URL } from "../http";
import { ErrorResponse } from "../models/errorResponse/ErrorResponse";

export default class UserStore {
    isAuth = false;
    role = "";

    constructor() {
        makeAutoObservable(this);
    }

    setAuth(bool: boolean) {
        this.isAuth = bool;
    }

    setRole(role: string) {
        this.role = role;
    }

    async login(role: string, password: string) : Promise<ErrorResponse> {
        try {
            const response = await AuthService.login(role, password);
            localStorage.setItem("token", response.data.access_token);
            localStorage.setItem("refresh_token", response.data.refresh_token);
            this.setAuth(true);
            this.setRole(response.data.role)
            return {message: "", code: 200} 
        } catch (err) {
            if (axios.isAxiosError(err))  {
                return {message: err.response?.data, code: err.response?.status}
            }
            return {message: "", code: 400} 
        }
    }

    async logout() {
        try {
            localStorage.removeItem("token");
            this.setAuth(false);
            this.setRole("")
            localStorage.clear()
        } catch (e) {
            console.log(e)
        }
    }

    async checkAuth() {
        try {
            const response = await axios.post<AuthResponse>(`${API_URL}/auth/refresh`, { "token": localStorage.getItem("refresh_token") }, { withCredentials: true });
            localStorage.setItem("token", response.data.access_token);
            localStorage.setItem("refresh_token", response.data.refresh_token);
            this.setAuth(true);
            this.setRole(response.data.role)
        } catch (e) {
            this.logout();
        }
    }
}