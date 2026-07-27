<script setup>
/**
 * HelloWorld.vue — Example Interactive Component
 *
 * This component demonstrates the Wails Go ↔ JavaScript bridge by:
 * 1. Taking user input from a text field
 * 2. Calling the Go backend method `Greet()` via the auto-generated bridge
 * 3. Displaying the returned greeting message
 */

import { reactive } from 'vue';
import { Greet } from '../../wailsjs/go/main/App';

/**
 * Reactive application data.
 *
 * @property {string} name       - The user-provided name input
 * @property {string} resultText - The greeting message returned from the Go backend
 */
const data = reactive({
  name: '',
  resultText: 'Please enter your name below 👇',
});

/**
 * Calls the Go backend's Greet() method with the current name value,
 * then updates the result text with the server's response.
 *
 * This function is triggered when the user clicks the "Greet" button.
 */
function greet() {
  Greet(data.name).then((result) => {
    data.resultText = result;
  });
}
</script>

<template>
  <main>
    <!-- Display area for the greeting result returned from the Go backend -->
    <div id="result" class="result">{{ data.resultText }}</div>

    <!-- Input area with name field and greet button -->
    <div id="input" class="input-box">
      <input
        id="name"
        v-model="data.name"
        autocomplete="off"
        class="input"
        type="text"
      />
      <button class="btn" @click="greet">Greet</button>
    </div>
  </main>
</template>

<style scoped>
.result {
  height: 20px;
  line-height: 20px;
  margin: 1.5rem auto;
}

.input-box .btn {
  width: 60px;
  height: 30px;
  line-height: 30px;
  border-radius: 3px;
  border: none;
  margin: 0 0 0 20px;
  padding: 0 8px;
  cursor: pointer;
}

.input-box .btn:hover {
  background-image: linear-gradient(to top, #cfd9df 0%, #e2ebf0 100%);
  color: #333333;
}

.input-box .input {
  border: none;
  border-radius: 3px;
  outline: none;
  height: 30px;
  line-height: 30px;
  padding: 0 10px;
  background-color: rgba(240, 240, 240, 1);
  -webkit-font-smoothing: antialiased;
}

.input-box .input:hover,
.input-box .input:focus {
  border: none;
  background-color: rgba(255, 255, 255, 1);
}
</style>
