// API response type
export interface ApiResponse<T = any> {
  data?: T;
  error?: string;
  message?: string;
  status: number;
}

// User types
export interface User {
  id: string;
  email: string;
  createdAt: string;
}

export interface AuthToken {
  accessToken: string;
  refreshToken: string;
  expiresAt: string;
}

// Expense types
export interface Expense {
  id: string;
  userId: string;
  amount: number;
  description: string;
  date: string;
  mainCategory: MainCategory;
  subCategory: SubCategory;
  createdAt: string;
  updatedAt: string;
}

export enum MainCategory {
  ChennaiHouse = 'Chennai House',
  BangaloreHouse = 'Bangalore House',
  Self = 'Self',
  Savings = 'Savings',
}

export enum SubCategory {
  Food = 'Food',
  Entertainment = 'Entertainment',
  Education = 'Education',
  Travel = 'Travel',
  Misc = 'Misc',
}

export interface Income {
  id: string;
  userId: string;
  amount: number;
  month: string;
  source: string;
  createdAt: string;
}

export interface MonthlyBreakdown {
  month: string;
  totalIncome: number;
  totalExpenses: number;
  savings: number;
  categoryBreakdown: Record<MainCategory, number>;
  subCategoryBreakdown: Record<SubCategory, number>;
}

// Task types
export interface Task {
  id: string;
  userId: string;
  title: string;
  description: string;
  dueDate?: string;
  priority: Priority;
  status: TaskStatus;
  labels: string[];
  isRecurring: boolean;
  createdAt: string;
  updatedAt: string;
}

export enum Priority {
  Low = 0,
  Medium = 1,
  High = 2,
}

export enum TaskStatus {
  Pending = 0,
  InProgress = 1,
  Done = 2,
}

export interface Reminder {
  id: string;
  taskId: string;
  time: string;
  triggered: boolean;
  createdAt: string;
}

// Note types
export interface Note {
  id: string;
  userId: string;
  title: string;
  content: string;
  tags: string[];
  isFavorite: boolean;
  attachments: Attachment[];
  createdAt: string;
  updatedAt: string;
}

export interface Attachment {
  id: string;
  type: AttachmentType;
  url: string;
  metadata: Record<string, any>;
  createdAt: string;
}

export enum AttachmentType {
  Link = 'link',
  Image = 'image',
}