package handler

import (
	"miya/merchants/api/internal/logic"
	"miya/merchants/api/internal/svc"
	"miya/merchants/api/internal/types"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func getMerchantHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.MerchantReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewGetMerchantLogic(r.Context(), svcCtx)
		resp, err := l.GetMerchant(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
