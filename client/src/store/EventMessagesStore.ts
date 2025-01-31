import EventMessage from '@/classes/EventMessage';

class S extends PF.H.BaseStore<EventMessage> {
  constructor() {
    super(EventMessage, 'event-messages');
  }
}

const store: S = new S();
export default store;
