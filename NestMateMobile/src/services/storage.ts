import PouchDB from 'pouchdb-react-native';
import AsyncStorageAdapter from 'pouchdb-adapter-asyncstorage';
import { Expense, Task, Note } from '../types';

// Initialize PouchDB with AsyncStorage adapter
PouchDB.plugin(AsyncStorageAdapter);

// Database instances
const expensesDB = new PouchDB('expenses');
const tasksDB = new PouchDB('tasks');
const notesDB = new PouchDB('notes');

// Expense storage service
export const ExpenseStorage = {
  async saveExpense(expense: Expense): Promise<string> {
    try {
      const response = await expensesDB.put({
        _id: expense.id,
        ...expense,
      });
      return response.id;
    } catch (error) {
      console.error('Error saving expense:', error);
      throw error;
    }
  },

  async getExpense(id: string): Promise<Expense | null> {
    try {
      const doc = await expensesDB.get(id);
      return doc as unknown as Expense;
    } catch (error) {
      console.error('Error getting expense:', error);
      return null;
    }
  },

  async getAllExpenses(): Promise<Expense[]> {
    try {
      const result = await expensesDB.allDocs({ include_docs: true });
      return result.rows.map((row: any) => row.doc) as unknown as Expense[];
    } catch (error) {
      console.error('Error getting all expenses:', error);
      return [];
    }
  },

  async deleteExpense(id: string): Promise<boolean> {
    try {
      const doc = await expensesDB.get(id);
      await expensesDB.remove(doc);
      return true;
    } catch (error) {
      console.error('Error deleting expense:', error);
      return false;
    }
  },
};

// Task storage service
export const TaskStorage = {
  async saveTask(task: Task): Promise<string> {
    try {
      const response = await tasksDB.put({
        _id: task.id,
        ...task,
      });
      return response.id;
    } catch (error) {
      console.error('Error saving task:', error);
      throw error;
    }
  },

  async getTask(id: string): Promise<Task | null> {
    try {
      const doc = await tasksDB.get(id);
      return doc as unknown as Task;
    } catch (error) {
      console.error('Error getting task:', error);
      return null;
    }
  },

  async getAllTasks(): Promise<Task[]> {
    try {
      const result = await tasksDB.allDocs({ include_docs: true });
      return result.rows.map((row: any) => row.doc) as unknown as Task[];
    } catch (error) {
      console.error('Error getting all tasks:', error);
      return [];
    }
  },

  async deleteTask(id: string): Promise<boolean> {
    try {
      const doc = await tasksDB.get(id);
      await tasksDB.remove(doc);
      return true;
    } catch (error) {
      console.error('Error deleting task:', error);
      return false;
    }
  },
};

// Note storage service
export const NoteStorage = {
  async saveNote(note: Note): Promise<string> {
    try {
      const response = await notesDB.put({
        _id: note.id,
        ...note,
      });
      return response.id;
    } catch (error) {
      console.error('Error saving note:', error);
      throw error;
    }
  },

  async getNote(id: string): Promise<Note | null> {
    try {
      const doc = await notesDB.get(id);
      return doc as unknown as Note;
    } catch (error) {
      console.error('Error getting note:', error);
      return null;
    }
  },

  async getAllNotes(): Promise<Note[]> {
    try {
      const result = await notesDB.allDocs({ include_docs: true });
      return result.rows.map((row: any) => row.doc) as unknown as Note[];
    } catch (error) {
      console.error('Error getting all notes:', error);
      return [];
    }
  },

  async deleteNote(id: string): Promise<boolean> {
    try {
      const doc = await notesDB.get(id);
      await notesDB.remove(doc);
      return true;
    } catch (error) {
      console.error('Error deleting note:', error);
      return false;
    }
  },
};

// Sync service
export const SyncService = {
  async syncExpenses(remoteURL: string): Promise<void> {
    try {
      await expensesDB.sync(new PouchDB(remoteURL + '/expenses'));
    } catch (error) {
      console.error('Error syncing expenses:', error);
      throw error;
    }
  },

  async syncTasks(remoteURL: string): Promise<void> {
    try {
      await tasksDB.sync(new PouchDB(remoteURL + '/tasks'));
    } catch (error) {
      console.error('Error syncing tasks:', error);
      throw error;
    }
  },

  async syncNotes(remoteURL: string): Promise<void> {
    try {
      await notesDB.sync(new PouchDB(remoteURL + '/notes'));
    } catch (error) {
      console.error('Error syncing notes:', error);
      throw error;
    }
  },
};