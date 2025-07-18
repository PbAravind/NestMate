// Type declarations for modules without TypeScript definitions
declare module 'pouchdb-react-native';
declare module 'pouchdb-adapter-asyncstorage';

// Fix for React types
declare module 'react' {
  import React from 'react';
  export = React;
}