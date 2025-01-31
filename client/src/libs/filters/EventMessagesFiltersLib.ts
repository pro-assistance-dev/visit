import EventMessage from '@/classes/EventMessage';

const EventMessagesFiltersLib = (() => {
  function byEventId(id?: string): FilterModel {
    const filterModel = FilterModel.Create(EventMessage, PF.H.Classes.GetPropertyName(EventMessage).eventId, DataTypes.String);
    filterModel.value1 = id;
    return filterModel;
  }

  return {
    byEventId,
  };
})();

export default EventMessagesFiltersLib;
