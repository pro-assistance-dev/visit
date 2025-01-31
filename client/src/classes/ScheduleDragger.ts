export default class ScheduleDragger<Schedule extends IWithId, DragItem> {
  private dragger = new Dragger();
  stretchedDown = false;

  savedStart = '';
  savedEnd = '';
  savedSchedule?: Schedule;

  time = '';
  schedule?: Schedule;
  dragItem?: ITimePeriod & IWithId;

  start(e: DragEvent, item: ITimePeriod & IWithId, schedule: Schedule) {
    this.savedStart = item.startTime;
    this.savedEnd = item.endTime;
    this.savedSchedule = schedule;
    this.dragItem = item;

    this.dragger.start(e);
  }

  stretchDown(e: DragEvent, item: ITimePeriod) {
    this.savedStart = item.startTime;
    this.savedEnd = item.endTime;
    this.dragItem = item;
    this.stretchedDown = true;
    this.dragger.start(e);
    console.log(item.startTime);
  }

  end() {
    this.stretchedDown = false;
    this.dragger.end();
  }

  active(): boolean {
    return this.dragger.isDragged();
  }

  setSchedule(schedule: Schedule): void {
    if (schedule.id === this.schedule?.id) {
      return;
    }
    this.schedule = schedule;
  }
}
