<template>
  <div class="footer">
    <div class="footer-content main-container">
      <div class="footer-content-top">
        <div>
          <p>Подпишитесь на наши новости</p>
          <div class="subscribe">
            <input type="text" placeholder="Введите Ваш адрес электронной почты" />
            <button class="main-btn">Отправить</button>
          </div>
        </div>
      </div>
      <div class="footer-content-bottom">© ПРО содействие 2019</div>
    </div>
  </div>
</template>

<script lang="ts" setup></script>

<style scoped>
.footer {
  color: white;
  overflow: hidden;
  background-color: #3d1819;
  height: 400px;

  &-content {
    padding: 24px 15px 0;
    &-top {
      padding-top: 48px;
      display: flex;
      justify-content: space-between;

      .subscribe {
        display: flex;
        align-items: stretch;
        padding-bottom: 5px;
        border-bottom: 1px solid white;
        width: 100%;
        input {
          width: 300px;
          border-radius: 0;
        }
        input:focus {
          outline: none !important;
          outline: 0;
          outline-color: initial;
          outline-style: initial;
          outline-width: 0px;
          box-shadow: 0 0 0 0.2rem rgb(0 123 255 / 25%);
        }
        button {
          margin-left: 10px;
          text-transform: uppercase;
          &:hover {
            cursor: pointer;
            color: #fff;
            background-color: #6c757d;
          }
        }
        input,
        button {
          font-size: 16px;
          font-weight: 400;
          line-height: 1.5;
          color: #495057;
          padding: 6px 12px;
          border: none;
          background: transparent;
        }
      }
    }
    &-bottom {
      margin: 48px 15px 0;
      color: #6c757d;
      font-size: 16px;
    }
  }
  ul {
    list-style-type: none;
    li {
      margin: 0 20px 20px;
    }
  }

  a {
    color: white;
    font-size: 16px;
    font-weight: bold;
    line-height: 1.5;
    text-decoration: none;
  }
  a:hover {
    cursor: pointer;
    text-decoration: underline;
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
