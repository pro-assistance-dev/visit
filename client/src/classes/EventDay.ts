export default class EventDay {
  id?: string;
  date = new Date(0);
  description = '';
  eventId? = '';

  constructor(i?: Event) {
    PF.H.Classes.BuildClass(this, i);
  }

  static Create(eventId: string, date = new Date()): EventDay {
    const item = new EventDay();
    item.id = PF.H.Id.UUID();
    item.eventId = eventId;
    item.date = date;
    return item;
  }
}
