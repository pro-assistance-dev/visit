<template>
  <div v-if="event.eventDays.length" class="program-table">
    <h1 id="program" :style="{ color: '#262338' }">Программа. 13.11.24</h1>
    <div v-for="schedule in event.schedules.filter((s) => s.perfoms.length > 0)" :key="schedule.id">
      <PTable>
        <template #thead>
          <tr>
            <th :style="{ width: '100px' }">Время</th>
            <th :style="{ width: '30%', textAlign: 'left' }">Тема</th>
            <th :style="{ width: 'calc(70% - 100px)', textAlign: 'left' }">Лектор</th>
          </tr>
        </template>
        <template #tbody>
          <tr v-for="perfom in schedule.perfoms" :key="perfom.id">
            <td :style="{ textAlign: 'center' }">{{ Time.GetPeriod(perfom.startTime, perfom.endTime) }}</td>
            <td>{{ perfom.name }}</td>
            <td>
              <div v-for="speaker in perfom.speakers" :key="speaker.id">
                {{ speaker.human.getFullName() }},
                {{ speaker.getFullRegalias() }}
              </div>
            </td>
          </tr>
        </template>
      </PTable>
    </div>
  </div>
</template>

<script setup lang="ts">
const event = Store.Events.Item;
</script>
