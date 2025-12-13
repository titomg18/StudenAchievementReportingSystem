package service

import (
	repo "StudenAchievementReportingSystem/app/repository/mongodb"
)

type AchievementService struct {
	achievementRepo repo.AchievementRepository
}