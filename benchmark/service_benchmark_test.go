package benchmark

import (
	"context"
	"database/sql"
	"emailn/internal/domain/campaign"
	"emailn/internal/infrastructure/database"
	"emailn/internal/infrastructure/db"
	"testing"
)

func BenchmarkGetCampaigns(b *testing.B) {
	conn, _ := db.OpenConn()
	defer func(DB *sql.DB) { _ = DB.Close() }(conn.DB)

	ctx := context.Background()
	repo := database.NewCampaignRepository(ctx, conn.DB)

	service1 := &campaign.ServiceImpl{Repository: repo} // Crie a instância do seu service com o repositório real

	b.ResetTimer() // Reseta o timer do benchmark para não incluir o tempo de setup

	for i := 0; i < b.N; i++ {
		_, _ = service1.GetBy(20)
	}
}
