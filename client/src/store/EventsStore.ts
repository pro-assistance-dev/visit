import Event from '@/classes/Event';
import { ChatMessage } from '@/modules/chats/classes';

class S extends PF.H.BaseStore<Event> {
  constructor() {
    super(Event, 'events');
  }

  async GetBySlug(slug?: string) {
    this.Set(await PF.H.HttpClient.Get<Event>({ query: this.getUrl('by-slug/' + slug) }));
  }
  async CreateMessage(m: ChatMessage<unknown>) {
    PF.H.HttpClient.Post({ query: this.getUrl('message'), payload: m });
    // this.Set(await PF.H.HttpClient.Get<Event>({ query: this.getUrl('by-slug/' + slug) }));
  }
}

const store: S = new S();
export default store;
