<template>Главная</template>

<script lang="ts" setup></script>

<style scoped>
.colored {
  padding: 200px 0;
  background: linear-gradient(135deg, #f25a2a 18%, #f67f1e 41%, rgba(255, 222, 0, 0.51) 100%);
}

.image-container {
  padding: 60px 0;
}
.ellipse-icon {
  margin-right: 20px;
  height: 86px;
  background: #fff;
  border-radius: 100%;
  text-align: center;
  display: flex;
  flex-direction: column;
  justify-content: center;
  padding: 0 22px;
  img {
    width: 42px;
    vertical-align: middle;
    border-style: none;
    overflow-clip-margin: content-box;
    overflow: clip;
    box-sizing: border-box;
  }
}
.services {
  margin-top: 50px;
  column-count: 2;
  &-item {
    width: 60%;
    display: flex;
    align-items: center;
    margin: 0 auto;
    padding-top: 50px;
    font-size: 16px;
    font-weight: bold;
  }
}
.contacts {
  &-content {
    display: flex;
    margin: 48px 0;
  }
  &-content-left {
    width: 60%;

    &-info {
      display: flex;
      align-items: center;
      img {
        margin-right: 24px;
      }
    }

    &-links {
      margin: 16px 0 0 48px;
      img {
        margin: 10px;
      }
    }
  }
  &-content-right {
    width: 40%;
    display: flex;
    flex-direction: column;
    input {
      font-size: 16px;
      font-weight: 400;
      line-height: 1.5;
      color: #495057;
      padding: 6px 12px;
      border: none;
      background: transparent;
      border-bottom: 1px solid black;
      margin-bottom: 15px;
    }
    input:focus {
      outline: none !important;
      outline: 0;
      outline-color: initial;
      outline-style: initial;
      outline-width: 0px;
      box-shadow: 0 0 0 0.2rem rgb(0 123 255 / 25%);
    }
  }
}

@media screen and (max-width: 768px) {
  .services {
    margin-top: 20px;
    column-count: 1;
    padding: 0 20px;
    &-item {
      min-width: 280px;
      max-width: 500px;
      display: flex;
      align-items: center;
      margin: 0 auto;
      padding-top: 20px;
      font-size: 16px;
      font-weight: bold;
    }
  }

  .ul {
    display: block;
  }

  .contacts {
    &-content {
      display: block;
      margin: 20px 0;
    }
    &-content-left {
      min-width: 280px;
      max-width: 280px;
      margin: 0 auto;
    }
    &-content-right {
      min-width: 280px;
      max-width: 600px;
      margin: 0 auto;
    }
  }

  .image-container {
    padding: 60px 0;
  }
}

.main-btn {
  padding: 5px 15px;
  color: white;
  font-size: 16px;
  line-height: 1.5;
  border-radius: 30px;
  background-color: #3d1819;
  text-transform: uppercase;
  cursor: pointer;
  border: none;
}
</style>
