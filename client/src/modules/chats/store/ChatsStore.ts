import BaseStore from '@/core/helpers/BaseStore';
import Chat from '@/modules/chats/classes/Chat';

class S<UserT> extends BaseStore<Chat<UserT>> {
  constructor() {
    super(Chat, 'chats');
  }
}

const store: S<unknown> = new S();
export default store;
