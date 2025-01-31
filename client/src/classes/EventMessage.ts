import User from './User';

export default class EventMessage {
  id?: string;
  message = '';
  createdOn = new Date();
  user = new User();
  userId?: string;
  eventId?: string;

  constructor(i?: EventMessage) {
    PF.H.Classes.BuildClass(this, i);
  }

  static Create(eventId?: string, userId?: string, message?: string): EventMessage {
    const date = new Date();
    date.setTime(date.getTime() + 3 * 60 * 60 * 1000);
    return new EventMessage({ eventId: eventId, userId: userId, message: message?.trim(), createdOn: date } as EventMessage);
  }
  static GetClassName(): string {
    return 'eventMessage';
  }
}
