<template>
  <div class="chat">
    <div class="chat-body" id="chat-body" ref="chat-body" v-if="messages">
      <div
        class="chat-body-message-container"
        :class="[{ incoming: !user || message.user.id !== user.id }, { outgoing: user && message.user.id === user.id }]"
        v-for="message of messages"
        :key="message.id"
      >
        <div class="chat-body-avatar-container">
          <el-avatar class="chat-body-avatar" :src="`https://api.dicebear.com/7.x/initials/svg?seed=${message.user.firstName}`"></el-avatar>
        </div>
        <div class="chat-body-message">
          <div class="chat-body-message-header">
            <div>
              <span class="chat-body-message-header-name"> {{ message.user.lastName }}&nbsp; </span>
              <span class="chat-body-message-header-name">
                {{ message.user.firstName }}
              </span>
            </div>
            <!--                      <span class="msg_time">{{ $moment(message.created_on).format('DD.MM.YYYY HH:mm:ss') }}</span>-->
            <span class="chat-body-message-header-time">
              <!-- {{ moment.utc(message.createdOn).format('HH:mm') }} -->
            </span>
          </div>
          <div class="chat-body-message-body">
            {{ message.message }}
          </div>
        </div>
      </div>
    </div>
    <div class="chat-footer">
      <div class="chat-footer-message">
        <input
          @keyup.enter="sendMessage"
          type="text"
          v-model="newMessage"
          placeholder="Введите сообщение"
          class="chat-footer-message-input"
        />
        <button @click="sendMessage" class="chat-footer-message-button">Отправить</button>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import EventMessage from '@/classes/EventMessage';
import EventMessagesSortsLib from '@/libs/sorts/EventMessagesSortsLib';

const event = Store.Events.Item;
const messages = Store.EventMessages.Items;
const auth = PF.Auth;
const user = PF.Auth.Status.User;

const newMessage: Ref<string> = ref('');
const chatBody = ref();
// const f = new FilterQuery();
let i: number | undefined = undefined;

const ftsp = new PF.FTSP();
ftsp.setSortModel(EventMessagesSortsLib.byCreatedAt());
if (ftsp.p) {
  ftsp.p.limit = 10000;
}
// ftsp.replaceF(EventMessagesFiltersLib.byEventId(event.value.id));

const loadChat = async () => {
  await Store.EventMessages.FTSP({ ftps: ftsp });
  // messages.value.sort((a, b) => (a.createdOn > b.createdOn ? 1 : b.createdOn > a.createdOn ? -1 : 0));
  // await Provider.router.replace({ query: {} });
};

const sendMessage = async () => {
  console.log('newMessage.value', newMessage.value);
  if (newMessage.value === '') {
    return;
  }
  await Store.EventMessages.Create(EventMessage.Create(event.id, auth.value.user.getId(), newMessage.value));
  newMessage.value = '';
  await loadChat();
  document.getElementById('chat-body')?.scrollTo({ top: 999999999999 });
};

const cl = () => {
  clearInterval(i);
  i = undefined;
};
onMounted(async () => {
  // if (event.value.id) {
  //   f.filterModels.push(EventMessagesFiltersLib.byEventId(event.value.id));
  // }
  // const sort = EventMessagesSortsLib.byCreatedOn(Orders.Asc);
  // f.sortModelsonBeforeUnmount.push(sort);
  window.addEventListener('beforeunload', cl);

  // await Provider.store.dispatch('usersEventsActivities/ping');
  // await loadChat();
  // i = window.setInterval(async () => {
  //   await loadChat();
  //   await Provider.store.dispatch('usersEventsActivities/ping');
  // }, 50000);
  // document.getElementById('chat-body')?.scrollTo({ top: 999999999999 });
});

window.onbeforeunload = function () {
  clearInterval(i);
  i = undefined;
  ('Are you sure you want to close the window?');
};

onBeforeUnmount(() => {
  console.log('clearInterval');
  clearInterval(i);
  i = undefined;
});
</script>

<style lang="scss" scoped>
.incoming {
  .chat-body-message {
    background-color: #edffcb;
  }
  .chat-body-message::after {
    border-bottom: 10px solid #edffcb;
  }
}
.outgoing {
  .chat-body-message {
    background-color: #ffffc2;
  }
  .chat-body-message::after {
    border-bottom: 10px solid #ffffc2;
  }
}
.chat {
  box-sizing: border-box;
  height: 100%;
  display: flex;
  flex-direction: column;
  background-color: white;
  opacity: 0.9;
  min-width: 350px;
  max-width: 350px;
  padding: 10px 0 0 10px;
  font-size: 14px;
  &-body {
    width: 100%;
    overflow-y: scroll;
    height: 100%;

    &-message-container {
      margin-bottom: 5px;
      display: flex;
    }

    &-avatar-container {
      line-height: 30px;
      display: flex;
      align-items: flex-end;
    }
    &-avatar {
      height: 30px;
      width: 30px;
      line-height: 30px;
      vertical-align: bottom;
    }

    &-message {
      margin: auto 5px auto 10px;
      border-radius: 10px 10px 10px 0;
      padding: 10px;
      position: relative;
      min-width: 100px;
      font-family: Montserrat, sans-serif;
      &::after {
        content: '';
        position: absolute;
        border: 10px solid transparent;
        bottom: 0;
        left: -10px;
      }
      &-header {
        line-height: 20px;
        display: flex;
        align-items: center;
        justify-content: space-between;

        &-name {
          display: inline-block;
          vertical-align: middle;
          line-height: normal;
          font-weight: bold;
        }

        &-time {
          color: rgba(0, 0, 0, 0.5);
          margin-left: 10px;
          display: inline-block;
          vertical-align: middle;
          line-height: normal;
        }
      }

      &-body {
        font-weight: lighter;
      }
    }
  }

  &-footer {
    width: 100%;
    margin: 5px 0;

    &-message {
      z-index: 0;
      height: 30px;
      font-size: 0;
      font-weight: 500;
      display: flex;
      flex-wrap: nowrap;
      &-input {
        // height: 30px;
        width: 100%;
        border-color: #e3e9ef;
        font-size: 14px;
        font-family: 'Montserrat', sans-serif;
        border-right: none;
        border-top-left-radius: 10px;
        border-bottom-left-radius: 10px;
        padding: 10px;

        &:focus {
          outline: none;
        }

        &::placeholder {
          color: #aca6cc;
        }
      }
      &-button {
        font-size: 14px;
        font-family: 'Montserrat', sans-serif;
        font-weight: 500;
        color: white;
        height: 30px;
        border: none;
        background-color: #224af2;
        border-top-right-radius: 10px;
        border-bottom-right-radius: 10px;
        margin-right: 10px;

        &:hover {
          filter: brightness(130%);
          cursor: pointer;
        }
      }
    }
  }
}
@media screen and (max-width: 1230px) {
  .chat {
    max-width: 100%;
    height: 600px;
  }
}
@media screen and (max-width: 768px) {
  .chat {
    font-size: 12px;
    height: 500px;
  }
}
@media screen and (max-width: 480px) {
  .chat {
    height: 500px;
    &-footer-message-input {
      height: 25px;
      font-size: 12px;
    }
    &-footer-message-button {
      height: 25px;
      font-size: 12px;
    }
  }
}
::-webkit-scrollbar {
  display: block;
  width: 8px;
  height: 8px;
  background-color: rgba(245, 245, 245, 0.47);
}

::-webkit-scrollbar-track {
  border-radius: 10px;
  background-color: #f5f5f5;
}

::-webkit-scrollbar-thumb {
  height: 20px;
  border-radius: 10px;
  box-shadow: inset 0 0 6px rgba(0, 0, 0, 0.3);
  background-color: rgba(85, 85, 85, 0.25);
}

//input.chat-footer-message-input,
//textarea {
//  background-color: rgba(0, 0, 0, 1);
//  opacity: 0.5;
//  background-color: white;
//}
</style>
