import Perfom from './Perfom';
import Speaker from './Speaker';

export default class Session {
  id?: string;
  name = '';
  description = '';
  startTime = '';
  endTime = '';

  scheduleId?: string;

  @PF.H.Classes.GetClassConstructor(Perfom)
  perfoms: Perfom[] = [];

  @PF.H.Classes.GetClassConstructor(Speaker)
  chairs: Speaker[] = [];

  //
  show = true;
  supposedStart = '';
  supposedEnd = '';
  conflicted = false;
  moved = true;

  constructor(i?: Session) {
    PF.H.Classes.BuildClass(this, i);
    this.formatTime();
  }

  formatTime() {
    this.startTime = PF.H.Time.HmsToHm(this.startTime);
    this.endTime = PF.H.Time.HmsToHm(this.endTime);
  }

  static Create(): Session {
    const item = new Session();
    item.startTime = '9:00';
    item.endTime = '9:30';
    item.id = PF.H.Id.UUID();
    return item;
  }

  bindChair(item: Speaker): M2M {
    this.chairs.push(item);
    return this.createChairM2M(item.id as string);
  }

  unbindChair(id: string): M2M {
    PF.H.Classes.RemoveFromClassById(id, this.chairs, []);
    return this.createChairM2M(id as string);
  }

  createChairM2M(id: string): M2M {
    return M2M.CreateV1('sessionToSpeaker', { field: 'sessionId', value: this.id as string }, { field: 'speakerId', value: id });
  }

  setName(name: string) {
    this.name = name;
  }

  static GetClassName(): string {
    return 'session';
  }

  getDuration(): number {
    return Time.DiffM(this.startTime, this.endTime);
  }

  setPeriod(t1: string, t2?: string): void {
    const dur = this.getDuration();
    this.startTime = t1;
    this.endTime = t2 ?? Time.AddMtoHM(this.startTime, dur);
  }

  setSupposedPeriod(t1: string, t2?: string) {
    const dur = this.getDuration();
    this.supposedStart = t1;
    this.supposedEnd = t2 ?? Time.AddMtoHM(this.supposedStart, dur);
  }

  acceptPeriod(): void {
    this.startTime = this.supposedStart;
    this.endTime = this.supposedEnd;
  }

  getMacroInfo(): string {
    const n = this.name ?? '';
    const t = PF.H.Time.GetPeriod(this.startTime, this.endTime);
    // const p = this.chairId ? `Председатель: ${this.chair.human.getFullName()}` : '';
    return [t, n].filter((s) => s !== '').join('. ');
  }

  getMiddleInfo(): string {
    const n = this.name ?? '';
    const t = PF.H.Time.GetPeriod(this.startTime, this.endTime);
    // const p = this.chairId ? `${this.chair.human.getInitialsName()}` : '';
    return [t, n].filter((s) => s !== '').join('. ');
  }

  periodIsFree(t1: string, t2: string, perfomId?: string): boolean {
    return !this.perfoms
      .filter((s: Perfom) => s.id !== perfomId)
      .some((s: Perfom) => PF.H.Time.PeriodsIntersects(t1, t2, s.startTime, s.endTime));
  }
  allPerfomsInPeriod(): boolean {
    return this.perfoms.every((p: Perfom) => PF.H.Time.PeriodInPeriod(p.startTime, p.endTime, this.startTime, this.endTime));
  }
}
