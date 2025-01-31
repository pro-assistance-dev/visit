<template>
  <el-form>
    <LeftRightContainer left-width="280px">
      <template #left>
        <UploaderImage
          :file-info="speaker.human.photo"
          :height="280"
          :default-ratio="1"
          :emit-crop="true"
          @ratio="(e) => (element.ratio = e)"
          @crop="savePhoto"
          @removeFile="removePhoto"
        />
        <div class="asr"><AdminSpeakerRegalias :speaker="speaker" @update="update" /></div>
      </template>

      <template #right>
        <div class="container">
          <el-form-item label="Фамилия" prop="surname">
            <el-input
              v-model="speaker.human.surname"
              placeholder="Введите фамилию"
              formatter="firstLetterUpper"
              @blur="updateHuman()"
            ></el-input>
          </el-form-item>
          <el-form-item label="Имя" prop="name">
            <el-input v-model="speaker.human.name" placeholder="Введите имя" formatter="firstLetterUpper" @blur="updateHuman()"></el-input>
          </el-form-item>
          <el-form-item label="Отчество" prop="patronymic">
            <el-input
              v-model="speaker.human.patronymic"
              placeholder="Введите отчество"
              formatter="firstLetterUpper"
              @blur="updateHuman()"
            ></el-input>
          </el-form-item>
          <el-form-item label="Город" prop="city">
            <el-input v-model="speaker.city" placeholder="Город спикера" formatter="firstLetterUpper" @blur="update"></el-input>
          </el-form-item>
          <ContactForm :contact="speaker.human.contact" />
          <div>
            <div class="contact-container" :style="{ background: speaker.experiences.length ? '' : '#F9FAFB' }">
              <div class="bottom-buttons">
                <div class="ex-title" :style="{ color: !speaker.experiences.length ? '#c4c4c4' : '#303133' }">Опыт работы</div>
                <button class="admin-add" @click.prevent="addExperience">+ Добавить</button>
              </div>

              <div v-for="(experience, i) in speaker.experiences" :key="experience.id" class="contact-container-item">
                <button class="admin-del" @click.prevent="removeExperience(experience.id)">Удалить</button>
                <div class="list-number">
                  {{ i + 1 }}
                </div>
                <el-form>
                  <el-form-item label="Заведение" prop="surname">
                    <el-input
                      v-model="experience.place"
                      @blur="updateExp(experience)"
                      placeholder="Введите заведение"
                      formatter="firstLetterUpper"
                    ></el-input>
                  </el-form-item>
                  <el-form-item label="Подразделение" prop="name">
                    <el-input
                      v-model="experience.division"
                      @blur="updateExp(experience)"
                      placeholder="Введите подразделение"
                      formatter="firstLetterUpper"
                    ></el-input>
                  </el-form-item>
                  <el-form-item label="Должность" prop="patronymic">
                    <el-input
                      v-model="experience.position"
                      @blur="updateExp(experience)"
                      placeholder="Введите должность"
                      formatter="firstLetterUpper"
                    ></el-input>
                  </el-form-item>
                </el-form>
                <!-- <ExperienceConstructor :experience="experience" /> -->
              </div>
            </div>
          </div>
          <el-form-item label="Прочие должности" prop="description">
            <el-input
              type="textarea"
              v-model="speaker.position"
              placeholder="Введите описание"
              formatter="firstLetterUpper"
              @blur="update"
            />
          </el-form-item>
        </div>
      </template>
    </LeftRightContainer>
  </el-form>
</template>

<script lang="ts" setup>
import Speaker from '@/classes/Speaker';
import Experience from '@/classes/Experience';

const speaker: Speaker = SpeakersStore.Item;

const updateExp = async (experience: Experience) => {
  await ExperiencesStore.Update(experience);
};
const update = async () => {
  await SpeakersStore.Update(speaker);
};
const savePhoto = async () => {
  await FileInfosStore.Create(speaker.human.photo);
  speaker.human.photoId = speaker.human.photo.id;
  await updateHuman();
};
const removePhoto = async () => {
  // await Provider.store.dispatch('fileInfos/remove', speaker.human.photoId);
  speaker.human.photoId = undefined;
  speaker.human.photo = new FileInfo();
  await updateHuman();
};
const updateHuman = async () => {
  await HumansStore.Update(speaker.human);
};

const addExperience = async () => {
  const item = speaker.addExperience();
  await ExperiencesStore.Create(item);
};

const removeExperience = async (id: string) => {
  speaker.removeExperience(id);
  await ExperiencesStore.Remove(id);
};
onBeforeMount(async () => {
  await createContactInfo();
});

const createContactInfo = async () => {
  if (speaker.human.contact.id) {
    return;
  }
  speaker.human.initContact();
  await ContactsStore.Create(speaker.human.contact);
  await updateHuman();
};
</script>

<style lang="scss" scoped>
.light-title {
  color: #a1a8bd;
  display: flex;
  align-items: center;
}

.upper {
  text-transform: uppercase;
}

.flex-center {
  display: flex;
  align-items: center;
}

.contact-icon {
  font-size: 25px;
  margin-right: 5px;
}

.el-tag {
  font-size: 14px;
}

.contact-container-item {
  position: relative;
  width: calc(100% - 42px);
  margin: 0px 10px 20px 10px;
  border: 1px solid #c3c3c3;
  border-radius: 5px;
  padding: 12px 10px;
  background: #f9fafb;
}

.admin-add {
  border: none;
  background: inherit;
  color: #00b5a4;
  transition: 0.3s;
  cursor: pointer;
}

.admin-add:hover {
  color: #009e8f;
}

.admin-del {
  top: 23px;
  right: 36px;
  border: none;
  background: inherit;
  color: #a3a9be;
  transition: 0.3s;
  cursor: pointer;
}

.admin-del:hover {
  color: #b63414;
}

.bottom-buttons {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding-right: 15px;
}

.contact-container {
  box-sizing: border-box;
  width: 100%;
  position: relative;
  padding: 0 10px;
  border: 1px solid #c3c3c3;
  border-radius: 5px;
  margin: 0px 0px 20px 0px;
  background: #dff2f8;
}

.asr {
  max-width: 280px;
  background: #ffffff;
  padding: 10px;
  margin: 0 auto;
  box-sizing: border-box;
  margin-top: 10px;
  border-radius: 5px;
}

.save-button {
  width: 100%;
  border-radius: 5px;
  height: 42px;
  color: #006bb4;
  background: #dff2f8;
  margin: 2px 10px 0 0;
  font-size: 14px;
}

.del-button {
  width: 100%;
  border-radius: 5px;
  height: 42px;
  color: #ff4c68;
  background: #ecc7c7;
  margin: 2px 10px 0 0;
  font-size: 14px;
}

.container {
  box-sizing: border-box;
  width: 100%;
}

.register-tag {
  &:hover {
    cursor: pointer;
    border-width: 2px;
  }
}

:deep(.el-upload--picture-card) {
  width: 150px;
  font-size: 50px;
}

:deep(.el-upload--picture-card i) {
  font-size: 50px;
  color: #00b5a4;
  padding: 0 114px;
}

.left-title {
  font-size: 14px;
  font-weight: bold;
  color: #343d5c;
  margin-top: 28px;
  margin-left: 10px;
  text-transform: uppercase;
  display: flex;
  justify-content: left;
  align-items: center;
}

.left-info {
  font-size: 12px;
  color: #343d5c;
  margin-top: 3px;
  margin-left: 10px;
}

.info-title {
  width: 100%;
  box-sizing: border-box;
  padding: 10px;
  border: 1px solid #b83616;
  display: flex;
  justify-content: space-between;
  align-items: center;
  border-radius: 5px;
  margin-top: 40px;
}

.el-select {
  width: 100%;
}

:deep(.el-date-editor.el-input, .el-date-editor.el-input__inner) {
  width: 100%;
}

:deep(.el-form-item) {
  display: block;
  margin-bottom: 16px;
}

:deep(.el-input__inner) {
  height: 40px;
  width: 100%;
  display: flex;
  font-family: Comfortaa, Arial, Helvetica, sans-serif;
  font-size: 15px;
  color: #343d5c;
}

:deep(.el-input__inner::placeholder) {
  color: #b0a4c0;
}

:deep(.el-input__icon) {
  color: #343d5c;
}

:deep(.el-form-item__label) {
  color: #b0a4c0;
  padding: 0 !important;
  text-transform: uppercase;
  margin-left: 5px;
  font-size: 14px;
  margin-bottom: 6px;
}

:deep(.el-input-number__increase) {
  border-radius: 0;
}

:deep(.el-input-number__decrease) {
  border-radius: 0;
}

:deep(.el-textarea__inner) {
  font-family: Comfortaa, Arial, Helvetica, sans-serif;
  font-size: 15px;
  color: #343d5c;
}

.button-add {
  border: none;
  background: inherit;
  color: #00b5a4;
  transition: 0.3s;
  cursor: pointer;
}

.button-add:hover {
  color: #009e8f;
}

.button-del {
  position: absolute;
  top: 13px;
  right: 36px;
  border: none;
  background: inherit;
  color: #a3a9be;
  transition: 0.3s;
  cursor: pointer;
}

.button-del:hover {
  color: #b63414;
}

.bottom-buttons {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding-right: 15px;
}

.list-number {
  position: absolute;
  top: 10px;
  right: 10px;
  width: 22px;
  height: 22px;
  display: flex;
  align-items: center;
  justify-content: center;
  color: #ffffff;
  background: #1979cf;
  border-radius: 20px;
}

.ex-title {
  font-family: 'Open Sans', sans-serif;
  font-size: 14px;
  color: #c4c4c4;
  margin: 10px;
  color: #b0a4c0;
  margin-left: 5px;
  font-size: 14px;
}

.place-item {
  width: calc(100% - 22px);
  position: relative;
  padding: 0 10px;
  border: 1px solid #c3c3c3;
  border-radius: 5px;
  margin: 0px;
  background: #dff2f8;
}

.place-item-item {
  position: relative;
  width: calc(100% - 42px);
  margin: 0px 10px 20px 10px;
  border: 1px solid #c3c3c3;
  border-radius: 5px;
  padding: 12px 10px;
  background: #f9fafb;
}

.session-edit {
  position: absolute;
  top: 20px;
  right: 20px;
  border: 1px solid #343e5c;
  border-radius: 5px;
  width: 100%;
  background: #ffffff;
  padding: 10px;
  z-index: 5;
}

.icon-alarm {
  width: 40px;
  height: 40px;
  fill: #b83616;
  margin: 0 10px 0 0;
}

@media (max-width: 600px) {
  .container {
    margin-left: 10px;
  }
}
</style>
