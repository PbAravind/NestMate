import axios, { AxiosInstance } from 'axios';
import AsyncStorage from '@react-native-async-storage/async-storage';

const API_BASE_URL = __DEV__ 
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
      async (config) => {
        const token = await AsyncStorage.getItem('authToken');
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
      (response) => response,
      async (error) => {
        if (error.response?.status === 401) {
          // Token expired, redirect to login
          await AsyncStorage.removeItem('authToken');
          // Navigation logic will be implemented later
        }
        return Promise.reject(error);
      }
    );
  }

  // Auth endpoints
  async register(email: string, password: string) {
    return this.api.post('/auth/register', { email, password });
  }

  async login(email: string, password: string) {
    return this.api.post('/auth/login', { email, password });
  }

  async logout() {
    return this.api.post('/auth/logout');
  }

  // Expense endpoints
  async getExpenses(params?: any) {
    return this.api.get('/expenses', { params });
  }

  async createExpense(expense: any) {
    return this.api.post('/expenses', expense);
  }

  async updateExpense(id: string, expense: any) {
    return this.api.put(`/expenses/${id}`, expense);
  }

  async deleteExpense(id: string) {
    return this.api.delete(`/expenses/${id}`);
  }

  // Task endpoints
  async getTasks(params?: any) {
    return this.api.get('/tasks', { params });
  }

  async createTask(task: any) {
    return this.api.post('/tasks', task);
  }

  async updateTask(id: string, task: any) {
    return this.api.put(`/tasks/${id}`, task);
  }

  async deleteTask(id: string) {
    return this.api.delete(`/tasks/${id}`);
  }

  // Notes endpoints
  async getNotes(params?: any) {
    return this.api.get('/notes', { params });
  }

  async createNote(note: any) {
    return this.api.post('/notes', note);
  }

  async updateNote(id: string, note: any) {
    return this.api.put(`/notes/${id}`, note);
  }

  async deleteNote(id: string) {
    return this.api.delete(`/notes/${id}`);
  }

  async searchNotes(query: string) {
    return this.api.get('/notes/search', { params: { q: query } });
  }
}

export default new ApiService();