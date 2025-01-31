import BannersStore from './BannersStore';
import EventDaysStore from './EventDaysStore';
import EventMessagesStore from './EventMessagesStore';
import EventsStore from './EventsStore';
import ExperiencesStore from './ExperiencesStore';
import PerfomsStore from './PerfomsStore';
import PlacesStore from './PlacesStore';
import SchedulesStore from './SchedulesStore';
import SessionsStore from './SessionsStore';
import SpeakersStore from './SpeakersStore';
import SponsorsStore from './SponsorsStore';
import UsersStore from './UsersStore';

export abstract class Store {
  static Banners = BannersStore;
  static EventDays = EventDaysStore;
  static EventMessages = EventMessagesStore;
  static Events = EventsStore;
  static Experiences = ExperiencesStore;
  static Perfoms = PerfomsStore;
  static Places = PlacesStore;
  static Schedules = SchedulesStore;
  static Sessions = SessionsStore;
  static Speakers = SpeakersStore;
  static Sponsors = SponsorsStore;
  static Users = UsersStore;
}
