import Perfom from '@/classes/Perfom';
import Session from '@/classes/Session';

import ScheduleDragger from './ScheduleDragger';

export default class Schedule {
  id?: string;
  name = '';
  description = '';
  date = new Date();

  eventDayId?: string;
  placeId?: string;
  eventId?: string;

  @PF.H.Classes.GetClassConstructor(Perfom)
  perfoms: Perfom[] = [];

  @PF.H.Classes.GetClassConstructor(Session)
  sessions: Session[] = [];
  //

  constructor(i?: Schedule) {
    PF.H.Classes.BuildClass(this, i);
  }

  addPerfom(): Perfom {
    const item = Perfom.Create();
    item.scheduleId = this.id;
    this.perfoms.push(item);
    return item;
  }

  findNextPerfom(perfom: Perfom): Perfom | undefined {
    this.perfoms.sort((s1: Perfom, s2: Perfom) => Time.HMtoM(s1.startTime) - Time.HMtoM(s2.startTime));
    const nextPerfom = this.perfoms.find((p: Perfom) => Time.Gt(p.startTime, perfom.startTime));
    return nextPerfom;
  }

  findPrevPerfom(perfom: Perfom): Perfom | undefined {
    const c = this.perfoms.filter((p: Perfom) => p.show);
    c.sort((s1: Perfom, s2: Perfom) => Time.HMtoM(s1.startTime) - Time.HMtoM(s2.startTime));
    const i = c.findIndex((p: Perfom) => p.id === perfom.id);
    return c[i - 1];
  }

  findPrevSession(session: Session): Session | undefined {
    const c = this.sessions.filter((p: Session) => p.show);
    c.sort((s1: Session, s2: Session) => Time.HMtoM(s1.startTime) - Time.HMtoM(s2.startTime));
    const i = c.findIndex((p: Session) => p.id === session.id);
    return c[i - 1];
  }

  removePerfom(id?: string): void {
    PF.H.Classes.RemoveFromClassById(id, this.perfoms);
  }

  static Create(eventId?: string, eventDayId?: string, placeId?: string): Schedule {
    const item = new Schedule();
    item.id = PF.H.Id.UUID();
    item.placeId = placeId;
    item.eventDayId = eventDayId;
    item.eventId = eventId;
    return item;
  }

  addSession(): Session {
    const item = Session.Create();
    item.scheduleId = this.id;
    item.startTime = Arrays.GetLast(this.sessions)?.endTime ?? '9:00';
    item.endTime = Time.AddMtoHM(item.startTime, 30);
    this.sessions.push(item);
    return item;
  }

  getPerfomByStartTime(startTime: string): Perfom | undefined {
    return this.perfoms.find((p: Perfom) => Time.Eq(p.startTime, startTime));
  }

  removeSession(id?: string): void {
    PF.H.Classes.RemoveFromClassById(id, this.sessions, []);
  }

  lastItemEndTime(): string {
    return Arrays.GetLast(this.sessions).endTime ?? '9:00';
  }

  // getPerfomsBySessionId(sessionId?: string): Perfom[] {
  //   return this.perfoms.filter((p: Perfom) => p.sessionId === sessionId);
  // }

  findPerfom(id?: string): Perfom | undefined {
    return this.perfoms.find((p: Perfom) => p.id === id);
  }

  // PERFOM MOVE
  resolvePerfomsConflicts(perfom: Perfom, scheduler: Scheduler): Perfom[] {
    const problemPerfoms = this.perfoms.filter((p: Perfom) => {
      return !p.moved && p.id !== perfom.id && Time.PeriodsIntersects(perfom.supposedStart, perfom.supposedEnd, p.startTime, p.endTime);
    });
    console.log('problem:', problemPerfoms);
    const resolvedPerfoms: Perfom[] = [];
    problemPerfoms.forEach((p: Perfom) => {
      if (Time.Gt(p.startTime, perfom.supposedStart)) {
        p.setSupposedPeriod(perfom.supposedEnd);
        if (Time.GtOrEq(p.supposedEnd, scheduler.endHM) && p.getDuration() > 5) {
          p.supposedEnd = scheduler.endHM;
        }
      } else if (Time.Gt(perfom.supposedStart, p.startTime)) {
        const mDiff = Time.DiffM(perfom.supposedStart, p.endTime);
        p.setSupposedPeriod(Time.AddMtoHM(p.startTime, -mDiff));
        if (Time.GtOrEq(scheduler.startHM, p.supposedStart) && p.getDuration() > 5) {
          p.supposedStart = scheduler.startHM;
        }
      } else if (Time.Eq(perfom.supposedStart, p.startTime)) {
        p.setSupposedPeriod(perfom.supposedEnd);
        console.log(p, perfom);
      }

      if (!scheduler.periodIsCorrect(p.supposedStart, p.supposedEnd)) {
        p.conflicted = true;
      }
      p.moved = true;
      resolvedPerfoms.push(...this.resolvePerfomsConflicts(p, scheduler));
      scheduler.sortByStartTime(this.perfoms);
      this.perfoms.sort((s1: Perfom, s2: Perfom) => Time.HMtoM(s1.supposedStart) - Time.HMtoM(s2.supposedEnd));
      resolvedPerfoms.push(p);
    });
    return resolvedPerfoms;
  }

  //movePlan
  //
  movePerfom(perfom: Perfom, scheduler: Scheduler, dragger: ScheduleDragger<Schedule, Perfom>): Perfom[] {
    let movedPerfoms: Perfom[] = [];

    if (dragger.stretchedDown) {
      perfom.setSupposedPeriod(perfom.startTime, dragger.time);
    } else {
      perfom.setSupposedPeriod(dragger.time, undefined);
    }
    if (!scheduler.periodIsCorrect(perfom.supposedStart, perfom.supposedEnd)) {
      perfom.conflicted = true;
    }
    movedPerfoms = this.resolvePerfomsConflicts(perfom, scheduler);
    movedPerfoms.push(perfom);

    const confl = movedPerfoms.some((p: Perfom) => {
      return p.conflicted === true;
    });
    // if (movedPerfoms.some((p: Perfom) => !!p.conflicted)) {
    // console.log('someConflicted');
    // movedPerfoms = [];
    // } else {
    //   console.log('movedPerfomsLine');
    movedPerfoms.forEach((p: Perfom) => {
      p.acceptPeriod();
      p.conflicted = false;
      console.log(p);
    });
    // }
    this.perfoms.forEach((p: Perfom) => {
      p.conflicted = false;
      p.moved = false;
    });
    return movedPerfoms;
  }

  // SESSION MOVE
  resolveSessionsConflicts(perfom: Session, scheduler: Scheduler): Session[] {
    const problemPerfoms = this.sessions.filter((p: Session) => {
      return !p.moved && p.id !== perfom.id && Time.PeriodsIntersects(perfom.supposedStart, perfom.supposedEnd, p.startTime, p.endTime);
    });
    const resolvedPerfoms: Session[] = [];

    problemPerfoms.forEach((p: Session) => {
      if (Time.Gt(p.startTime, perfom.supposedStart)) {
        p.setSupposedPeriod(perfom.supposedEnd);
        if (Time.GtOrEq(p.supposedEnd, scheduler.endHM) && p.getDuration() > 5) {
          p.supposedEnd = scheduler.endHM;
        }
      } else if (Time.Gt(perfom.supposedStart, p.startTime)) {
        const mDiff = Time.DiffM(perfom.supposedStart, p.endTime);
        p.setSupposedPeriod(Time.AddMtoHM(p.startTime, -mDiff));
        if (Time.GtOrEq(scheduler.startHM, p.supposedStart) && p.getDuration() > 5) {
          p.supposedStart = scheduler.startHM;
        }
      } else if (Time.Eq(perfom.supposedStart, p.startTime)) {
        p.setSupposedPeriod(perfom.supposedEnd);
        console.log(p, perfom);
      }

      if (!scheduler.periodIsCorrect(p.supposedStart, p.supposedEnd)) {
        p.conflicted = true;
      }
      p.moved = true;
      resolvedPerfoms.push(...this.resolveSessionsConflicts(p, scheduler));
      scheduler.sortByStartTime(this.perfoms);
      this.perfoms.sort((s1: Perfom, s2: Perfom) => Time.HMtoM(s1.supposedStart) - Time.HMtoM(s2.supposedEnd));
      resolvedPerfoms.push(p);
    });
    return resolvedPerfoms;
  }

  moveSession(session: Session, scheduler: Scheduler, dragger: ScheduleDragger<Schedule, Perfom>): Session[] {
    let movedPerfoms: Session[] = [];

    if (dragger.stretchedDown) {
      session.setSupposedPeriod(session.startTime, dragger.time);
    } else {
      session.setSupposedPeriod(dragger.time, undefined);
    }

    if (!scheduler.periodIsCorrect(session.supposedStart, session.supposedEnd)) {
      session.conflicted = true;
    }
    session.moved = true;
    movedPerfoms = this.resolveSessionsConflicts(session, scheduler);
    movedPerfoms.push(session);

    // if (movedPerfoms.some((p: Session) => p.conflicted)) {
    //   movedPerfoms = [];
    // } else {
    movedPerfoms.forEach((p: Session) => {
      p.acceptPeriod();
      p.conflicted = false;
    });
    // }
    this.sessions.forEach((p: Session) => {
      p.conflicted = false;
      p.moved = true;
    });
    return movedPerfoms;
  }
}
