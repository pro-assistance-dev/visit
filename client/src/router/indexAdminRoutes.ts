import AdminEventsRoutes from '@/router/AdminEventsRoutes';

import AdminPerfomsRoutes from './AdminPerfomsRoutes';
import AdminPlacesRoutes from './AdminPlacesRoutes';
import AdminSpeakersRoutes from './AdminSpeakersRoutes';
import AdminSponsorsRoutes from './AdminSponsorsRoutes';

export default [...AdminEventsRoutes, ...AdminPlacesRoutes, ...AdminSpeakersRoutes, ...AdminPerfomsRoutes, ...AdminSponsorsRoutes];
