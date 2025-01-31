import Place from '@/classes/Place';
import EventDay from './EventDay';
import Perfom from './Perfom';
import Schedule from './Schedule';
import Session from './Session';
import Sponsor from './Sponsor';
import Chat from '@/modules/chats/classes/Chat';
import User from '@/classes/User';

export type EventNewDatesResult = {
  newSchedules: Schedule[];
  newEventDays: EventDay[];
  daysForDelete: EventDay[];
  schedulesForDelete: Schedule[];
  perfomsForUpdate: Perfom[];
};

export default class Event {
  id?: string;
  name = '';
  slug = '';
  type = '';
  organizer = '';
  description = '';

  chat = new Chat<User>();
  chatId?: string;

  image = new PF.C.FileInfo();
  imageId?: string;

  @PF.H.Classes.GetClassConstructor(EventDay)
  eventDays: EventDay[] = [];

  @PF.H.Classes.GetClassConstructor(Perfom)
  perfoms: Perfom[] = [];

  @PF.H.Classes.GetClassConstructor(Session)
  sessions: Session[] = [];

  @PF.H.Classes.GetClassConstructor(Place)
  places: Place[] = [];

  @PF.H.Classes.GetClassConstructor(Schedule)
  schedules: Schedule[] = [];

  @PF.H.Classes.GetClassConstructor(Event)
  events: Event[] = [];

  @PF.H.Classes.GetClassConstructor(Sponsor)
  sponsors: Sponsor[] = [];

  startDragSchedule?: Schedule;
  draggedPerfomId?: string;

  constructor(i?: Event) {
    PF.H.Classes.BuildClass(this, i);
    this.chat = new Chat<User>(i?.chat, User);
  }

  init(): void {
    this.id = PF.H.Id.UUID();
    this.eventDays.push(EventDay.Create(this.id), EventDay.Create(this.id));
  }

  getPeriod(): string {
    const d1 = Arrays.GetFirst(this.eventDays);
    const d2 = Arrays.GetLast(this.eventDays);
    return DatesFormatter.GetPeriod(d1?.date, d2?.date);
  }

  addSchedule(placeId?: string, eventDayId?: string): Schedule {
    const item = Schedule.Create(this.id, eventDayId, placeId);
    this.schedules.push(item);
    return item;
  }

  getSchedule(eventDayId: string, placeId: string): Schedule {
    return this.schedules.find((s: Schedule) => s.placeId === placeId && s.eventDayId === eventDayId) as Schedule;
  }

  getScheduleById(id: string): Schedule {
    return this.schedules.find((s: Schedule) => s.id === id) as Schedule;
  }

  setName(name: string) {
    this.name = name;
  }

  addPerfom(): Perfom {
    const item = Perfom.Create();
    item.eventId = this.id;
    this.perfoms.push(item);
    return item;
  }

  removePerfom(id?: string): void {
    PF.H.Classes.RemoveFromClassById(id, this.perfoms);
  }

  // addEventDay(): EventDay {
  // const item = EventDay.Create(this.id);
  // this.eventDays.push(item);
  // return item;
  // }

  static GetClassName(): string {
    return 'event';
  }

  bindPlace(place: Place): M2M {
    this.places.push(place);
    return this.createPlaceM2M(place.id as string);
  }

  unbindPlace(id: string): M2M {
    PF.H.Classes.RemoveFromClassById(id, this.places, []);
    return this.createPlaceM2M(id as string);
  }

  createPlaceM2M(placeId: string): M2M {
    return M2M.CreateV1('eventToPlace', { field: 'eventId', value: this.id as string }, { field: 'placeId', value: placeId });
  }

  placeExists(id: string): boolean {
    return !!this.places.find((p: Place) => p.id === id);
  }
  findDay(date: Date): EventDay | undefined {
    return this.eventDays.find((d: EventDay) => d.date === date);
  }
  findPlace(id: string): Place | undefined {
    return this.places.find((i: Place) => i.id === id);
  }
  addSchedulesInDay(e: EventDay): Schedule[] {
    const items: Schedule[] = [];
    this.places.forEach((p: Place) => {
      items.push(this.addSchedule(p.id, e.id));
    });
    return items;
  }

  dateExists(d: Date): boolean {
    return this.eventDays.some((e: EventDay) => {
      return Dates.Eq(e.date, d);
    });
  }

  setDates(start: Date, end: Date): EventNewDatesResult {
    const newDates = Dates.GetDatesFromPeriod(start, end);
    const newSchedules: Schedule[] = [];
    const newEventDays: EventDay[] = [];

    newDates.forEach((d: Date) => {
      if (!this.dateExists(d) && Dates.InPeriod(d, start, end)) {
        const newEventDay = EventDay.Create(this.id as string, d);
        newEventDay.date = d;
        newEventDays.push(newEventDay);
        newSchedules.push(...this.addSchedulesInDay(newEventDay));

        this.eventDays.push(newEventDay);
        this.schedules.push(...newSchedules);
      }
    });

    const daysForDelete: EventDay[] = [];
    const schedulesForDelete: Schedule[] = [];

    this.eventDays.forEach((e: EventDay, i: number) => {
      if (!Dates.InPeriod(e.date, start, end)) {
        daysForDelete.push(e);
        // PF.H.Classes.RemoveFromClassByIndex(i, this.eventDays);
        this.schedules.forEach((s: Schedule) => {
          if (s.eventDayId === e.id) {
            schedulesForDelete.push(s);
          }
        });
      }
    });

    this.eventDays = this.eventDays.filter((e: EventDay) => {
      return daysForDelete.every((d: EventDay) => d.id !== e.id);
    });

    const perfomsForUpdate: Perfom[] = [];

    schedulesForDelete.forEach((s: Schedule) => {
      s.perfoms.forEach((p: Perfom) => {
        p.scheduleId = undefined;
        p.eventId = this.id;
        this.perfoms.push(p);
        perfomsForUpdate.push(p);
      });
      s.perfoms = [];
    });

    return { newSchedules, newEventDays, daysForDelete, schedulesForDelete, perfomsForUpdate };
  }

  getStartDate(): Date {
    return Dates.GetDateOrToday(this.eventDays[0]);
  }

  getEndDate(): Date {
    return Dates.GetDateOrToday(this.eventDays[this.eventDays.length - 1]);
  }

  bindSponsor(place: Sponsor): M2M {
    this.sponsors.push(place);
    return this.createSponsorM2M(place.id as string);
  }

  unbindSponsor(id: string): M2M {
    PF.H.Classes.RemoveFromClassById(id, this.places, []);
    return this.createSponsorM2M(id as string);
  }

  createSponsorM2M(sponsorId: string): M2M {
    return M2M.CreateV1('eventToSponsor', { field: 'eventId', value: this.id as string }, { field: 'sponsorId', value: sponsorId });
  }

  moveSession(fromSchedule: Schedule, session: Session, toScheduleId?: string): void {
    const toSchedule = this.schedules.find((s: Schedule) => s.id === toScheduleId);
    if (!toSchedule || fromSchedule.id === toScheduleId) {
      return;
    }
    fromSchedule.removeSession(session.id);
    toSchedule.sessions.push(session);
    session.scheduleId = toSchedule.id;
  }

  copyPerfomToSchedule(fromSchedule: Schedule, perfom: Perfom, toScheduleId?: string): void {
    const toSchedule = this.schedules.find((s: Schedule) => s.id === toScheduleId);
    if (!toSchedule || fromSchedule.id === toScheduleId) {
      return;
    }

    const findedPerfom = toSchedule.findPerfom(perfom.id);
    if (findedPerfom) {
      findedPerfom.show = true;
      perfom.show = false;
    } else {
      const c = new Perfom(perfom);
      perfom.show = false;
      toSchedule.perfoms.push(c);
      c.scheduleId = toSchedule.id;
    }
  }

  unbindPerfomsFromPlace(placeId: string): unknown {
    const perfomsForUpdate: Perfom[] = [];
    const schedulesForDelete: Schedule[] = [];
    this.schedules.forEach((s: Schedule) => {
      if (s.placeId === placeId) {
        s.perfoms.forEach((p: Perfom) => {
          p.scheduleId = undefined;
          p.eventId = this.id;
          this.perfoms.push(p);
          perfomsForUpdate.push(p);
        });
        s.perfoms = [];
        schedulesForDelete.push(s);
      }
    });
    return { perfomsForUpdate, schedulesForDelete };
  }

  removeDuplicatesPerfoms(perfomId?: string, scheduleId?: string): void {
    console.log(perfomId, scheduleId);

    this.schedules.forEach((s: Schedule) => {
      s.perfoms = s.perfoms.filter((p: Perfom) => p.id !== perfomId || (p.id === perfomId && scheduleId === p.scheduleId));
    });
  }
}
