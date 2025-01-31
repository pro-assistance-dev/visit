<template>
  <GridContainer margin="0" grid-template-columns="repeat(auto-fit, minmax(220px, 1fr))">
    <InfoItemSelectButtons title="учёная степень" :selected="speaker.academicDegree" :list="degrees" @select="setAcademicDegree" />
    <InfoItemSelectButtons title="учёное звание" :selected="speaker.academicRank" :list="ranks" @select="setAcademicRank" />
    <InfoItemSelectButtons title="академия наук" :selected="speaker.academicRelation" :list="relations" @select="setAcademicRelation" />
  </GridContainer>
</template>

<script setup lang="ts">
import Speaker from '@/classes/Speaker';
import { AcademicDegree, AcademicRelation, AcademicRank } from '@/interfaces/AcademicEnums';

const props = defineProps({
  speaker: {
    type: Speaker as PropType<Speaker>,
    required: true,
  },
});
const degrees = Object.values(AcademicDegree);
const ranks = Object.values(AcademicRank);
const relations = Object.values(AcademicRelation);
const emit = defineEmits(['update']);

const setAcademicDegree = async (item: AcademicDegree) => {
  props.speaker.setAcademicDegree(item);
  emit('update');
};
const setAcademicRank = async (item: AcademicRank) => {
  props.speaker.setAcademicRank(item);
  emit('update');
};
const setAcademicRelation = async (item: AcademicRelation) => {
  props.speaker.setAcademicRelation(item);
  emit('update');
};
</script>

<style lang="scss" scoped></style>
