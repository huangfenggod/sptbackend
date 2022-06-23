package cron

import (
	"log"
	"sptbackend/sql"
)

type uidtable struct {
	Uid int `db:"uid"`
}

func CommuityDivid()  {
	var uids []uidtable
	sql.DB.Raw("select uid from user").Scan(&uids)
	for i:=0;i<len(uids);i++ {
	sumfather(uids[i].Uid)
	}
}

type usergrade struct{
	Uid int `db:"uid"`
	Pid int `db:"pid"`
	Grade int `db:"grade"`
	TodayPget int64 `db:"today_pget"`
}

func sumfather(id int)  {
	var uself usergrade
	sql.DB.Raw("select uid,pid,grade ,today_pget from user where uid = ?",id).Scan(&uself)

	sum := selectSon(uself)
	log.Printf("team reward uid:%d,amount: %d",id,sum)
	if sum>0 {
		sql.DB.Exec("update user set  cashable = cashable + ?,total_reward = total_reward + ? ,today_reward =? where uid=?",sum,sum,sum,id)
	}
}

//func getGradeRate(grade int)int64  {
//	var rate int64
//	switch grade{
//	case 4:
//		rate = 18
//	case 3:
//		rate = 15
//	case 2:
//		rate = 10
//	default:
//		rate =0
//	}
//	return rate
//}

func selectSon(user usergrade) int64  {
	sum :=int64(0)
	var rate int64

	var uson []usergrade
	sql.DB.Raw("select uid,pid,grade ,today_pget from user where pid = ? ",user.Uid).Scan(&uson)
	switch user.Grade {
	case 4:
		if len(uson)>0 {
			for i := 0; i < len(uson); i++ {
				rate = 18
				if user.Grade < uson[i].Grade {
					continue
				}else if user.Grade == uson[i].Grade{
					sum +=uson[i].TodayPget/100
					continue
				}
				if uson[i].Grade==3 {
					rate = 3
				}else if uson[i].Grade==2 {
					rate = 8
				}
				sum += uson[i].TodayPget * rate / 100
				var uson1 []usergrade
				sql.DB.Raw("select uid,pid,grade ,today_pget from user where pid = ? ", uson[i].Uid).Scan(&uson1)
				if len(uson1) > 0 {
					for a := 0; a < len(uson1); a++ {
						if user.Grade < uson1[a].Grade {
							continue
						}else if user.Grade == uson1[a].Grade{
							sum +=uson1[a].TodayPget/100
							continue
						}
						rate1 :=rate
						if uson1[a].Grade==3 ||rate==3{
							rate1 =3
						}else if  rate==8 ||uson1[i].Grade==2 {
							rate1 =8
						}
						sum += uson1[a].TodayPget * rate1 / 100
						var uson2 []usergrade
						sql.DB.Raw("select uid,pid,grade ,today_pget from user where pid = ? ", uson1[a].Uid).Scan(&uson2)
						if len(uson2) > 0 {
							for b := 0; b < len(uson2); b++ {
								if user.Grade < uson2[b].Grade {
									continue
								}else if user.Grade == uson2[b].Grade{
									sum +=uson2[b].TodayPget/100
									continue
								}
								rate2 :=rate1
								if uson2[b].Grade==3 ||rate==3{
									rate2 =3
								}else if  uson2[b].Grade==2 ||rate==8{
									rate2 =8
								}
								sum += uson2[b].TodayPget * rate2 / 100
								var uson3 []usergrade
								sql.DB.Raw("select uid,pid,grade ,today_pget from user where pid = ? ", uson2[b].Uid).Scan(&uson3)
								if len(uson3) > 0 {
									for c := 0; c < len(uson3); c++ {
										if user.Grade < uson3[c].Grade {
											continue
										}else if user.Grade == uson3[c].Grade{
											sum +=uson3[c].TodayPget/100
											continue
										}
										rate3 :=rate2
										if uson3[c].Grade==3 ||rate==3{
											rate3 =3
										}else if  uson3[c].Grade==2 ||rate==8{
											rate3 =8
										}
										sum += uson3[c].TodayPget * rate3 / 100
										var uson4 []usergrade
										sql.DB.Raw("select uid,pid,grade ,today_pget from user where pid = ? ", uson3[c].Uid).Scan(&uson4)
										if len(uson4) > 0 {
											for d := 0; d < len(uson4); d++ {
												if user.Grade < uson4[d].Grade {
													continue
												}else if user.Grade == uson4[d].Grade{
													sum +=uson4[d].TodayPget/100
													continue
												}
												rate4 :=rate3
												if uson4[d].Grade==3 ||rate==3{
													rate4 =3
												}else if  uson4[d].Grade==2 ||rate==8{
													rate4 =8
												}
												sum += uson4[d].TodayPget * rate4 / 100
												var uson5 []usergrade
												sql.DB.Raw("select uid,pid,grade ,today_pget from user where pid = ? ", uson4[d].Uid).Scan(&uson5)
												if len(uson5) > 0 {
													for e := 0; e < len(uson5); e++ {
														if user.Grade < uson5[e].Grade {
															continue
														}else if user.Grade == uson5[e].Grade{
															sum +=uson5[e].TodayPget/100
															continue
														}
														rate5 :=rate4
														if uson5[e].Grade==3 ||rate==3{
															rate5 =3
														}else if  uson5[e].Grade==2 ||rate==8{
															rate5 =8
														}
														sum += uson5[e].TodayPget * rate5 / 100
														var uson6 []usergrade
														sql.DB.Raw("select uid,pid,grade ,today_pget from user where pid = ? ", uson5[e].Uid).Scan(&uson6)
														if len(uson6) > 0 {
															for f := 0; f < len(uson6); f++ {
																if user.Grade < uson6[f].Grade {
																	continue
																}else if user.Grade == uson6[f].Grade{
																	sum +=uson6[f].TodayPget/100
																	continue
																}
																rate6 :=rate5
																if uson6[f].Grade==3 ||rate==3{
																	rate6 =3
																}else if  uson6[f].Grade==2 ||rate==8{
																	rate6 =8
																}
																sum += uson6[f].TodayPget * rate6 / 100
																var uson7  []usergrade
																sql.DB.Raw("select uid,pid,grade ,today_pget from user where pid = ? ", uson6[f].Uid).Scan(&uson7)
																if len(uson7) > 0 {
																	for g := 0; g < len(uson7); g++ {
																		if user.Grade < uson7[g].Grade {
																			continue
																		}else if user.Grade == uson7[g].Grade{
																			sum +=uson7[g].TodayPget/100
																			continue
																		}
																		rate7 :=rate6
																		if uson7[g].Grade==3 ||rate==3{
																			rate7 =3
																		}else if  uson7[g].Grade==2 ||rate==8{
																			rate7 =8
																		}
																		sum += uson7[g].TodayPget * rate7 / 100

																		var uson8  []usergrade
																		sql.DB.Raw("select uid,pid,grade ,today_pget from user where pid = ? ", uson7[g].Uid).Scan(&uson8)
																		if len(uson8) > 0 {
																			for h := 0; h < len(uson8); h++ {
																				if user.Grade < uson8[h].Grade {
																					continue
																				} else if user.Grade == uson8[h].Grade {
																					sum += uson8[h].TodayPget / 100
																					continue
																				}
																				rate8 :=rate7
																				if uson8[h].Grade==3 ||rate==3{
																					rate8 =3
																				}else if  uson8[h].Grade==2 ||rate==8{
																					rate8 =8
																				}
																				sum += uson8[h].TodayPget * rate8 / 100

																				var uson9  []usergrade
																				sql.DB.Raw("select uid,pid,grade ,today_pget from user where pid = ? ", uson8[h].Uid).Scan(&uson9)
																				if len(uson9) > 0 {
																					for j := 0; j < len(uson9); j++ {
																						if user.Grade < uson9[j].Grade {
																							continue
																						} else if user.Grade == uson9[j].Grade {
																							sum += uson9[j].TodayPget / 100
																							continue
																						}
																						rate9 :=rate8
																						if uson9[j].Grade==3 ||rate==3{
																							rate9 =3
																						}else if  uson9[j].Grade==2 ||rate==8{
																							rate9 =8
																						}

																						sum += uson9[j].TodayPget * rate9 / 100
																						var uson10  []usergrade
																						sql.DB.Raw("select uid,pid,grade ,today_pget from user where pid = ? ", uson9[j].Uid).Scan(&uson10)
																						if len(uson10) > 0 {
																							for k := 0; k < len(uson9); k++ {
																								if user.Grade < uson10[k].Grade {
																									continue
																								} else if user.Grade == uson10[k].Grade {
																									sum += uson10[k].TodayPget / 100
																									continue
																								}
																								rate10 :=rate9
																								if uson10[k].Grade==3 ||rate==3{
																									rate10 =3
																								}else if  uson10[k].Grade==2 ||rate==8{
																									rate10 =8
																								}
																								sum += uson10[k].TodayPget * rate10 / 100

																							}
																						}
																					}
																				}
																			}
																		}
																	}
																}
															}
														}
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	case 3:
		if len(uson)>0 {
			for i := 0; i < len(uson); i++ {
				rate = 15
				if user.Grade < uson[i].Grade {
					continue
				}else if user.Grade == uson[i].Grade{
					sum +=uson[i].TodayPget/100
					continue
				}
				 if uson[i].Grade==2 {
					rate = 5
				}
				sum += uson[i].TodayPget * rate / 100
				var uson1 []usergrade
				sql.DB.Raw("select uid,pid,grade ,today_pget from user where pid = ? ", uson[i].Uid).Scan(&uson1)
				if len(uson1) > 0 {
					for a := 0; a < len(uson1); a++ {
						if user.Grade < uson1[a].Grade {
							continue
						}else if user.Grade == uson1[a].Grade{
							sum +=uson1[a].TodayPget/100
							continue
						}
						rate1 :=rate
						 if  rate==5 ||uson1[a].Grade==2 {
							rate1 =5
						}
						sum += uson1[a].TodayPget * rate1 / 100
						var uson2 []usergrade
						sql.DB.Raw("select uid,pid,grade ,today_pget from user where pid = ? ", uson1[a].Uid).Scan(&uson2)
						if len(uson2) > 0 {
							for b := 0; b < len(uson2); b++ {
								if user.Grade < uson2[b].Grade {
									continue
								}else if user.Grade == uson2[b].Grade{
									sum +=uson2[b].TodayPget/100
									continue
								}
								rate2 :=rate1
								if  uson2[b].Grade==2 ||rate==5{
									rate2 =5
								}
								sum += uson2[b].TodayPget * rate2 / 100
								var uson3 []usergrade
								sql.DB.Raw("select uid,pid,grade ,today_pget from user where pid = ? ", uson2[b].Uid).Scan(&uson3)
								if len(uson3) > 0 {
									for c := 0; c < len(uson3); c++ {
										if user.Grade < uson3[c].Grade {
											continue
										}else if user.Grade == uson3[c].Grade{
											sum +=uson3[c].TodayPget/100
											continue
										}
										rate3 :=rate2
										 if uson3[c].Grade==2 ||rate==5{
											rate3 =5
										}
										sum += uson3[c].TodayPget * rate3 / 100
										var uson4 []usergrade
										sql.DB.Raw("select uid,pid,grade ,today_pget from user where pid = ? ", uson3[c].Uid).Scan(&uson4)
										if len(uson4) > 0 {
											for d := 0; d < len(uson4); d++ {
												if user.Grade < uson4[d].Grade {
													continue
												}else if user.Grade == uson4[d].Grade{
													sum +=uson4[d].TodayPget/100
													continue
												}
												rate4 :=rate3
												if  uson4[d].Grade==2 ||rate==5{
													rate4 =5
												}
												sum += uson4[d].TodayPget * rate4 / 100
												var uson5 []usergrade
												sql.DB.Raw("select uid,pid,grade ,today_pget from user where pid = ? ", uson4[d].Uid).Scan(&uson5)
												if len(uson5) > 0 {
													for e := 0; e < len(uson5); e++ {
														if user.Grade < uson5[e].Grade {
															continue
														}else if user.Grade == uson5[e].Grade{
															sum +=uson5[e].TodayPget/100
															continue
														}
														rate5 :=rate4
														if  uson5[e].Grade==2 ||rate==5{
															rate5 =5
														}
														sum += uson5[e].TodayPget * rate5 / 100
														var uson6 []usergrade
														sql.DB.Raw("select uid,pid,grade ,today_pget from user where pid = ? ", uson5[e].Uid).Scan(&uson6)
														if len(uson6) > 0 {
															for f := 0; f < len(uson6); f++ {
																if user.Grade < uson6[f].Grade {
																	continue
																}else if user.Grade == uson6[f].Grade{
																	sum +=uson6[f].TodayPget/100
																	continue
																}
																rate6 :=rate5
																if  uson6[f].Grade==2 ||rate==5{
																	rate6 =5
																}
																sum += uson6[f].TodayPget * rate6 / 100
																var uson7  []usergrade
																sql.DB.Raw("select uid,pid,grade ,today_pget from user where pid = ? ", uson6[f].Uid).Scan(&uson7)
																if len(uson7) > 0 {
																	for g := 0; g < len(uson7); g++ {
																		if user.Grade < uson7[g].Grade {
																			continue
																		}else if user.Grade == uson7[g].Grade{
																			sum +=uson7[g].TodayPget/100
																			continue
																		}
																		rate7 :=rate6
																		if  uson7[g].Grade==2 ||rate==5{
																			rate7 =5
																		}
																		sum += uson7[g].TodayPget * rate7 / 100

																		var uson8  []usergrade
																		sql.DB.Raw("select uid,pid,grade ,today_pget from user where pid = ? ", uson7[g].Uid).Scan(&uson8)
																		if len(uson8) > 0 {
																			for h := 0; h < len(uson8); h++ {
																				if user.Grade < uson8[h].Grade {
																					continue
																				} else if user.Grade == uson8[h].Grade {
																					sum += uson8[h].TodayPget / 100
																					continue
																				}
																				rate8 :=rate7
																				if  uson8[h].Grade==2 ||rate==5{
																					rate8 =5
																				}
																				sum += uson8[h].TodayPget * rate8 / 100

																				var uson9  []usergrade
																				sql.DB.Raw("select uid,pid,grade ,today_pget from user where pid = ? ", uson8[h].Uid).Scan(&uson9)
																				if len(uson9) > 0 {
																					for j := 0; j < len(uson9); j++ {
																						if user.Grade < uson9[j].Grade {
																							continue
																						} else if user.Grade == uson9[j].Grade {
																							sum += uson9[j].TodayPget / 100
																							continue
																						}
																						rate9 :=rate8
																						if  uson9[j].Grade==2 ||rate==5{
																							rate9 =5
																						}

																						sum += uson9[j].TodayPget * rate9 / 100
																						var uson10  []usergrade
																						sql.DB.Raw("select uid,pid,grade ,today_pget from user where pid = ? ", uson9[j].Uid).Scan(&uson10)
																						if len(uson10) > 0 {
																							for k := 0; k < len(uson9); k++ {
																								if user.Grade < uson10[k].Grade {
																									continue
																								} else if user.Grade == uson10[k].Grade {
																									sum += uson10[k].TodayPget / 100
																									continue
																								}
																								rate10 :=rate9
																								if  uson10[k].Grade==2 ||rate==5{
																									rate10 =5
																								}
																								sum += uson10[k].TodayPget * rate10 / 100

																							}
																						}
																					}
																				}
																			}
																		}
																	}
																}
															}
														}
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	case 2:
		if len(uson)>0 {
			for i := 0; i < len(uson); i++ {
				rate = 10
				if user.Grade < uson[i].Grade {
					continue
				}else if user.Grade == uson[i].Grade{
					sum +=uson[i].TodayPget/100
					continue
				}

				sum += uson[i].TodayPget * rate / 100
				var uson1 []usergrade
				sql.DB.Raw("select uid,pid,grade ,today_pget from user where pid = ? ", uson[i].Uid).Scan(&uson1)
				if len(uson1) > 0 {
					for a := 0; a < len(uson1); a++ {
						if user.Grade < uson1[a].Grade {
							continue
						}else if user.Grade == uson1[a].Grade{
							sum +=uson1[a].TodayPget/100
							continue
						}

						sum += uson1[a].TodayPget * rate / 100
						var uson2 []usergrade
						sql.DB.Raw("select uid,pid,grade ,today_pget from user where pid = ? ", uson1[a].Uid).Scan(&uson2)
						if len(uson2) > 0 {
							for b := 0; b < len(uson2); b++ {
								if user.Grade < uson2[b].Grade {
									continue
								}else if user.Grade == uson2[b].Grade{
									sum +=uson2[b].TodayPget/100
									continue
								}

								sum += uson2[b].TodayPget * rate / 100
								var uson3 []usergrade
								sql.DB.Raw("select uid,pid,grade ,today_pget from user where pid = ? ", uson2[b].Uid).Scan(&uson3)
								if len(uson3) > 0 {
									for c := 0; c < len(uson3); c++ {
										if user.Grade < uson3[c].Grade {
											continue
										}else if user.Grade == uson3[c].Grade{
											sum +=uson3[c].TodayPget/100
											continue
										}

										sum += uson3[c].TodayPget * rate / 100
										var uson4 []usergrade
										sql.DB.Raw("select uid,pid,grade ,today_pget from user where pid = ? ", uson3[c].Uid).Scan(&uson4)
										if len(uson4) > 0 {
											for d := 0; d < len(uson4); d++ {
												if user.Grade < uson4[d].Grade {
													continue
												}else if user.Grade == uson4[d].Grade{
													sum +=uson4[d].TodayPget/100
													continue
												}

												sum += uson4[d].TodayPget * rate / 100
												var uson5 []usergrade
												sql.DB.Raw("select uid,pid,grade ,today_pget from user where pid = ? ", uson4[d].Uid).Scan(&uson5)
												if len(uson5) > 0 {
													for e := 0; e < len(uson5); e++ {
														if user.Grade < uson5[e].Grade {
															continue
														}else if user.Grade == uson5[e].Grade{
															sum +=uson5[e].TodayPget/100
															continue
														}

														sum += uson5[e].TodayPget * rate / 100
														var uson6 []usergrade
														sql.DB.Raw("select uid,pid,grade ,today_pget from user where pid = ? ", uson5[e].Uid).Scan(&uson6)
														if len(uson6) > 0 {
															for f := 0; f < len(uson6); f++ {
																if user.Grade < uson6[f].Grade {
																	continue
																}else if user.Grade == uson6[f].Grade{
																	sum +=uson6[f].TodayPget/100
																	continue
																}

																sum += uson6[f].TodayPget * rate / 100
																var uson7  []usergrade
																sql.DB.Raw("select uid,pid,grade ,today_pget from user where pid = ? ", uson6[f].Uid).Scan(&uson7)
																if len(uson7) > 0 {
																	for g := 0; g < len(uson7); g++ {
																		if user.Grade < uson7[g].Grade {
																			continue
																		}else if user.Grade == uson7[g].Grade{
																			sum +=uson7[g].TodayPget/100
																			continue
																		}

																		sum += uson7[g].TodayPget * rate / 100

																		var uson8  []usergrade
																		sql.DB.Raw("select uid,pid,grade ,today_pget from user where pid = ? ", uson7[g].Uid).Scan(&uson8)
																		if len(uson8) > 0 {
																			for h := 0; h < len(uson8); h++ {
																				if user.Grade < uson8[h].Grade {
																					continue
																				} else if user.Grade == uson8[h].Grade {
																					sum += uson8[h].TodayPget / 100
																					continue
																				}

																				sum += uson8[h].TodayPget * rate / 100

																				var uson9  []usergrade
																				sql.DB.Raw("select uid,pid,grade ,today_pget from user where pid = ? ", uson8[h].Uid).Scan(&uson9)
																				if len(uson9) > 0 {
																					for j := 0; j < len(uson9); j++ {
																						if user.Grade < uson9[j].Grade {
																							continue
																						} else if user.Grade == uson9[j].Grade {
																							sum += uson9[j].TodayPget / 100
																							continue
																						}


																						sum += uson9[j].TodayPget * rate / 100
																						var uson10  []usergrade
																						sql.DB.Raw("select uid,pid,grade ,today_pget from user where pid = ? ", uson9[j].Uid).Scan(&uson10)
																						if len(uson10) > 0 {
																							for k := 0; k < len(uson9); k++ {
																								if user.Grade < uson10[k].Grade {
																									continue
																								} else if user.Grade == uson10[k].Grade {
																									sum += uson10[k].TodayPget / 100
																									continue
																								}

																								sum += uson10[k].TodayPget * rate / 100

																							}
																						}
																					}
																				}
																			}
																		}
																	}
																}
															}
														}
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	default:
		return int64(0)

	}
return sum
}


