import Speaker from './Speaker';

export default class Perfom {
  id?: string;
  name = '';
  description = '';
  startTime = '';
  endTime = '';
  type = '';
  format = '';
  scheduleId?: string;
  eventId?: string;

  @PF.H.Classes.GetClassConstructor(Speaker)
  speakers: Speaker[] = [];
  //
  color = PF.H.Color.GetRandColor();
  show = true;
  moved = false;
  supposedStart = '';
  supposedEnd = '';
  conflicted = false;

  constructor(i?: Perfom) {
    PF.H.Classes.BuildClass(this, i);
    this.formatTime();
  }

  formatTime() {
    this.startTime = PF.H.Time.HmsToHm(this.startTime);
    this.endTime = PF.H.Time.HmsToHm(this.endTime);
  }

  static Create(): Perfom {
    const item = new Perfom();
    item.id = PF.H.Id.UUID();
    // item.name = "Новая локация"
    return item;
  }

  addSpeaker(): void {
    this.speakers.push(new Speaker());
  }

  speakerExists(speakerId: string): boolean {
    return !!this.speakers.find((s: Speaker) => s.id === speakerId);
  }

  setName(name: string) {
    this.name = name;
  }

  bindSpeaker(item: Speaker): M2M {
    this.speakers.push(item);
    return this.createSpeakerM2M(item.id as string);
  }

  unbindSpeaker(id: string): M2M {
    PF.H.Classes.RemoveFromClassById(id, this.speakers, []);
    return this.createSpeakerM2M(id as string);
  }
  createSpeakerM2M(id: string): M2M {
    return M2M.CreateV1('perfomToSpeaker', { field: 'perfomId', value: this.id as string }, { field: 'speakerId', value: id });
  }
  static GetClassName(): string {
    return 'perfom';
  }
  getDuration(): number {
    return Time.DiffM(this.startTime, this.endTime);
  }
  getTimeStep(divider: number): number {
    return this.getDuration() / divider;
  }

  setPeriod(t1: string, t2?: string): void {
    const dur = this.getDuration();
    this.startTime = t1;
    this.endTime = t2 ?? Time.AddMtoHM(this.startTime, dur);
  }

  acceptPeriod(): void {
    this.startTime = this.supposedStart;
    this.endTime = this.supposedEnd;
  }

  setSupposedPeriod(t1: string, t2?: string) {
    const dur = this.getDuration();
    this.supposedStart = t1;
    this.supposedEnd = t2 ?? Time.AddMtoHM(this.supposedStart, dur);
  }

  getMacroInfo(): string {
    const n = this.name ?? '';
    const t = Time.GetPeriod(this.startTime, this.endTime);
    const sNames = this.speakers.map((s: Speaker) => s.human.getFullName());
    const s = sNames.length > 0 ? 'Спикеры: ' + sNames.join(', ') : 'Спикеры отсутствуют ';
    return [t, n, s].filter((s) => s !== '').join('. ');
  }
}
