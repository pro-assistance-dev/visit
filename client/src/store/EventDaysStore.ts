import EventDay from '@/classes/EventDay';

class S extends PF.H.BaseStore<EventDay> {
  constructor() {
    super(EventDay, 'event-days');
  }
}

const store: S = new S();
export default store;
