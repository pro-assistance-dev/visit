import Event from '@/classes/Event';
import EventMessage from '@/classes/EventMessage';
const EventsSortsLib = (() => {
  function byCreatedAt(order?: Orders) {
    // return PF.C.SortModel.Create(
    //   EventMessage,
    //   PF.H.Classes.GetPropertyName(EventMessage).createdOn,
    //   order ? order : Orders.Asc,
    //   `По времеи ${order === Orders.Asc ? '(вверх)' : '(вниз)'}`,
    //   order === Orders.Desc ? false : true
    // );
  }
  return { byCreatedAt };
})();

export default EventsSortsLib;
