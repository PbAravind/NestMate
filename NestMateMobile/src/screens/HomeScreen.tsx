import React from 'react';
import {
  View,
  Text,
  StyleSheet,
  TouchableOpacity,
  ScrollView,
} from 'react-native';

const HomeScreen: React.FC = () => {
  return (
    <ScrollView style={styles.container}>
      <View style={styles.header}>
        <Text style={styles.title}>NestMate</Text>
        <Text style={styles.subtitle}>Your Personal Productivity Hub</Text>
      </View>

      <View style={styles.moduleGrid}>
        <TouchableOpacity style={styles.moduleCard}>
          <Text style={styles.moduleTitle}>Expenses</Text>
          <Text style={styles.moduleDescription}>Track your spending</Text>
        </TouchableOpacity>

        <TouchableOpacity style={styles.moduleCard}>
          <Text style={styles.moduleTitle}>Tasks</Text>
          <Text style={styles.moduleDescription}>Manage your to-dos</Text>
        </TouchableOpacity>

        <TouchableOpacity style={styles.moduleCard}>
          <Text style={styles.moduleTitle}>Notes</Text>
          <Text style={styles.moduleDescription}>Capture your thoughts</Text>
        </TouchableOpacity>

        <TouchableOpacity style={styles.moduleCard}>
          <Text style={styles.moduleTitle}>PDF Parser</Text>
          <Text style={styles.moduleDescription}>Import bank statements</Text>
        </TouchableOpacity>
      </View>
    </ScrollView>
  );
};

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: '#f5f5f5',
  },
  header: {
    padding: 20,
    backgroundColor: '#fff',
    marginBottom: 20,
  },
  title: {
    fontSize: 28,
    fontWeight: 'bold',
    color: '#333',
    marginBottom: 5,
  },
  subtitle: {
    fontSize: 16,
    color: '#666',
  },
  moduleGrid: {
    flexDirection: 'row',
    flexWrap: 'wrap',
    padding: 10,
    justifyContent: 'space-between',
  },
  moduleCard: {
    width: '48%',
    backgroundColor: '#fff',
    padding: 20,
    borderRadius: 10,
    marginBottom: 15,
    shadowColor: '#000',
    shadowOffset: {
      width: 0,
      height: 2,
    },
    shadowOpacity: 0.1,
    shadowRadius: 3.84,
    elevation: 5,
  },
  moduleTitle: {
    fontSize: 18,
    fontWeight: 'bold',
    color: '#333',
    marginBottom: 5,
  },
  moduleDescription: {
    fontSize: 14,
    color: '#666',
  },
});

export default HomeScreen;