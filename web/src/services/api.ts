import axios, { AxiosInstance, AxiosResponse } from 'axios';
import { ApiResponse } from '../types';

const API_BASE_URL = process.env.NODE_ENV === 'development' 
  ? 'http://localhost:8080/api/v1' 
  : 'https://your-production-api.com/api/v1';

class ApiService {
  private api: AxiosInstance;

  constructor() {
    this.api = axios.create({
      baseURL: API_BASE_URL,
      timeout: 10000,
      headers: {
        'Content-Type': 'application/json',
      },
    });

    // Request interceptor to add auth token
    this.api.interceptors.request.use(
      (config) => {
        const token = localStorage.getItem('authToken');
        if (token) {
          config.headers.Authorization = `Bearer ${token}`;
        }
        return config;
      },
      (error) => {
        return Promise.reject(error);
      }
    );

    // Response interceptor for error handling
    this.api.interceptors.response.use(
      (response: AxiosResponse) => response,
      (error) => {
        if (error.response?.status === 401) {
          // Token expired, redirect to login
          localStorage.removeItem('authToken');
          window.location.href = '/login';
        }
        return Promise.reject(error);
      }
    );
  }

  // Auth endpoints
  async register(email: string, password: string): Promise<ApiResponse<any>> {
    const response = await this.api.post('/auth/register', { email, password });
    return response.data;
  }

  async login(email: string, password: string): Promise<ApiResponse<any>> {
    const response = await this.api.post('/auth/login', { email, password });
    return response.data;
  }

  async logout(): Promise<ApiResponse<any>> {
    const response = await this.api.post('/auth/logout');
    return response.data;
  }

  // Expense endpoints
  async getExpenses(params?: any): Promise<ApiResponse<any>> {
    const response = await this.api.get('/expenses', { params });
    return response.data;
  }

  async createExpense(expense: any): Promise<ApiResponse<any>> {
    const response = await this.api.post('/expenses', expense);
    return response.data;
  }

  async updateExpense(id: string, expense: any): Promise<ApiResponse<any>> {
    const response = await this.api.put(`/expenses/${id}`, expense);
    return response.data;
  }

  async deleteExpense(id: string): Promise<ApiResponse<any>> {
    const response = await this.api.delete(`/expenses/${id}`);
    return response.data;
  }

  async getMonthlyBreakdown(month: string): Promise<ApiResponse<any>> {
    const response = await this.api.get('/expenses/breakdown', { params: { month } });
    return response.data;
  }

  // Task endpoints
  async getTasks(params?: any): Promise<ApiResponse<any>> {
    const response = await this.api.get('/tasks', { params });
    return response.data;
  }

  async createTask(task: any): Promise<ApiResponse<any>> {
    const response = await this.api.post('/tasks', task);
    return response.data;
  }

  async updateTask(id: string, task: any): Promise<ApiResponse<any>> {
    const response = await this.api.put(`/tasks/${id}`, task);
    return response.data;
  }

  async deleteTask(id: string): Promise<ApiResponse<any>> {
    const response = await this.api.delete(`/tasks/${id}`);
    return response.data;
  }

  async updateTaskStatus(id: string, status: string): Promise<ApiResponse<any>> {
    const response = await this.api.patch(`/tasks/${id}/status`, { status });
    return response.data;
  }

  // Notes endpoints
  async getNotes(params?: any): Promise<ApiResponse<any>> {
    const response = await this.api.get('/notes', { params });
    return response.data;
  }

  async createNote(note: any): Promise<ApiResponse<any>> {
    const response = await this.api.post('/notes', note);
    return response.data;
  }

  async updateNote(id: string, note: any): Promise<ApiResponse<any>> {
    const response = await this.api.put(`/notes/${id}`, note);
    return response.data;
  }

  async deleteNote(id: string): Promise<ApiResponse<any>> {
    const response = await this.api.delete(`/notes/${id}`);
    return response.data;
  }

  async searchNotes(query: string): Promise<ApiResponse<any>> {
    const response = await this.api.get('/notes/search', { params: { q: query } });
    return response.data;
  }
}

export default new ApiService();