import * as Keychain from 'react-native-keychain';
import AsyncStorage from '@react-native-async-storage/async-storage';
import api from './api';
import { AuthToken, User } from '../types';

const AUTH_TOKEN_KEY = 'authToken';
const REFRESH_TOKEN_KEY = 'refreshToken';
const USER_DATA_KEY = 'userData';

export const AuthService = {
  async login(email: string, password: string): Promise<User> {
    try {
      const response = await api.login(email, password);
      const { accessToken, refreshToken } = response.data as AuthToken;
      
      // Store tokens securely
      await Keychain.setGenericPassword(
        AUTH_TOKEN_KEY,
        accessToken,
        { service: AUTH_TOKEN_KEY }
      );
      
      await Keychain.setGenericPassword(
        REFRESH_TOKEN_KEY,
        refreshToken,
        { service: REFRESH_TOKEN_KEY }
      );
      
      // Store user data
      const userData = response.data.user as User;
      await AsyncStorage.setItem(USER_DATA_KEY, JSON.stringify(userData));
      
      return userData;
    } catch (error) {
      console.error('Login error:', error);
      throw error;
    }
  },
  
  async register(email: string, password: string): Promise<User> {
    try {
      const response = await api.register(email, password);
      return response.data as User;
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
    await Keychain.resetGenericPassword({ service: AUTH_TOKEN_KEY });
    await Keychain.resetGenericPassword({ service: REFRESH_TOKEN_KEY });
    await AsyncStorage.removeItem(USER_DATA_KEY);
  },
  
  async getAuthToken(): Promise<string | null> {
    try {
      const credentials = await Keychain.getGenericPassword({ service: AUTH_TOKEN_KEY });
      return credentials ? credentials.password : null;
    } catch (error) {
      console.error('Error getting auth token:', error);
      return null;
    }
  },
  
  async getRefreshToken(): Promise<string | null> {
    try {
      const credentials = await Keychain.getGenericPassword({ service: REFRESH_TOKEN_KEY });
      return credentials ? credentials.password : null;
    } catch (error) {
      console.error('Error getting refresh token:', error);
      return null;
    }
  },
  
  async getCurrentUser(): Promise<User | null> {
    try {
      const userData = await AsyncStorage.getItem(USER_DATA_KEY);
      return userData ? JSON.parse(userData) : null;
    } catch (error) {
      console.error('Error getting current user:', error);
      return null;
    }
  },
  
  async isAuthenticated(): Promise<boolean> {
    const token = await this.getAuthToken();
    return token !== null;
  },
};