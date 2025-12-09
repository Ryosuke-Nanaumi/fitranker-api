package training

//async getPersonalInfo(userId: number): Promise<PersonalUserInfo> {
//const user = await this.trainingRepository.getUserById(userId);
//
//if (!user) throw Error("user not found.");
//
//const [totalRecords, todaysRecords] = await Promise.all([
//this.trainingRepository.getPoint(user.id),
//this.trainingRepository.getPoint(user.id, this.date),
//]);
//
//const totalPoint = this.calcPoint(totalRecords);
//const todaysPoint = this.calcPoint(todaysRecords);
//
//return {
//...user,
//todaysPoint: todaysPoint,
//totalPoints: totalPoint,
//};
//}
//calcPoint(target: { amount: number; point: number }[]): number {
//return target.reduce((acc, cur) = > acc + cur.amount * cur.point, 0);
//}
//
//func (s *service) GetPersonalUser() (*PersonalUser, error) {
//	panic("TODO: implement DoSomething")
//	//user, err := s.re
//}
