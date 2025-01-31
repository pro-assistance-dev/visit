import Event from '@/classes/Event';
const EventsSortsLib = (() => {
  function byName(order?: Orders): SortModel {
    return SortModel.Create(
      Event,
      PF.H.Classes.GetPropertyName(Event).name,
      order ? order : Orders.Asc,
      `По ФИО ${order === Orders.Asc ? '(вверх)' : '(вниз)'}`,
      order === Orders.Desc ? false : true
    );
  }
  return { byName };
})();

export default EventsSortsLib;
