<script setup>
/**
 * App.vue — Root Application Component
 *
 * This is the top-level component of the application.
 * It renders a full-screen background image with a semi-transparent
 * white overlay for readability, and displays the app name
 * in the top-left corner.
 *
 * Future components will be added below the header area.
 */
</script>

<template>
  <!--
    Root container with background image and semi-transparent overlay.
    The .app-container establishes a full-viewport area with the
    background.webp image, while .overlay provides a white semi-transparent
    layer to reduce visual noise from the background.
  -->
  <div class="app-container">
    <!-- Semi-transparent white overlay to improve content readability -->
    <div class="overlay"></div>

    <!-- Header area: app name displayed in the top-left corner -->
    <header class="app-header">
      <h1 class="app-title">Yukumo Script</h1>
    </header>

    <!-- Main content area for future components -->
    <main class="app-content">
    </main>
  </div>
</template>

<style>
/* Root container: fills viewport and displays background image */
.app-container {
  position: relative;
  width: 100vw;
  height: 100vh;
  background-image: url('./assets/images/background.webp');
  background-size: cover;
  background-position: center;
  background-repeat: no-repeat;
  display: flex;
  flex-direction: column;
  overflow: hidden;
}

/* White semi-transparent overlay for readability */
.overlay {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(255, 255, 255, 0.7);
  pointer-events: none;
  z-index: 1;
}

/* Header positioned at the top-left */
.app-header {
  position: relative;
  z-index: 2;
  padding: 24px 32px;
  text-align: left;
}

/* App name styling */
.app-title {
  font-size: 28px;
  font-weight: 700;
  color: #2c3e50;
  margin: 0;
  letter-spacing: 1px;
}

/* Main content area for future components */
.app-content {
  position: relative;
  z-index: 2;
  flex: 1;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: flex-start;
  padding: 0 32px;
}
</style>