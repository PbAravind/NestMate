import api from './api';
import { AuthToken, User } from '../types';

const AUTH_TOKEN_KEY = 'authToken';
const REFRESH_TOKEN_KEY = 'refreshToken';
const USER_DATA_KEY = 'userData';

export const AuthService = {
  async login(email: string, password: string): Promise<User> {
    try {
      const response = await api.login(email, password);
      
      // The API returns an ApiResponse<any> where data contains the auth token and user
      if (response.data) {
        const authData = response.data as { token: AuthToken, user: User };
        const { token, user } = authData;
        
        // Store tokens in localStorage
        localStorage.setItem(AUTH_TOKEN_KEY, token.accessToken);
        localStorage.setItem(REFRESH_TOKEN_KEY, token.refreshToken);
        
        // Store user data
        localStorage.setItem(USER_DATA_KEY, JSON.stringify(user));
        
        return user;
      }
      
      throw new Error('Invalid response from server');
    } catch (error) {
      console.error('Login error:', error);
      throw error;
    }
  },
  
  async register(email: string, password: string): Promise<User> {
    try {
      const response = await api.register(email, password);
      
      // The API returns an ApiResponse<any> where data contains the user
      if (response.data && response.data.user) {
        return response.data.user as User;
      }
      
      throw new Error('Invalid response from server');
    } catch (error) {
      console.error('Registration error:', error);
      throw error;
    }
  },
  
  async logout(): Promise<void> {
    try {
      await api.logout();
    } catch (error) {
      console.error('Logout API error:', error);
      // Continue with local logout even if API fails
    }
    
    // Clear tokens and user data
    localStorage.removeItem(AUTH_TOKEN_KEY);
    localStorage.removeItem(REFRESH_TOKEN_KEY);
    localStorage.removeItem(USER_DATA_KEY);
  },
  
  getAuthToken(): string | null {
    return localStorage.getItem(AUTH_TOKEN_KEY);
  },
  
  getRefreshToken(): string | null {
    return localStorage.getItem(REFRESH_TOKEN_KEY);
  },
  
  getCurrentUser(): User | null {
    const userData = localStorage.getItem(USER_DATA_KEY);
    return userData ? JSON.parse(userData) : null;
  },
  
  isAuthenticated(): boolean {
    return this.getAuthToken() !== null;
  },
};