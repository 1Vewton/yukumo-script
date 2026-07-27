/**
 * main.js — Vue Application Entry Point
 *
 * This file bootstraps the Vue 3 application and mounts it to the DOM.
 * It imports the root component (App.vue) and global styles.
 */

import { createApp } from 'vue';
import App from './App.vue';
import './style.css';

// Create the Vue application instance and mount it to the #app element
createApp(App).mount('#app');
