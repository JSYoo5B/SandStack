package placement

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/JSYoo5B/SandStack/internal/api/respond"
	appplacement "github.com/JSYoo5B/SandStack/internal/app/placement"
)

func (h Handler) listAllocationCandidates(
	w http.ResponseWriter,
	r *http.Request,
) {
	candidates := h.service.ListAllocationCandidates(
		appplacement.AllocationCandidateQuery{
			Resources: parseRequestedResources(
				r.URL.Query().Get("resources"),
			),
		},
	)

	respond.JSON(w, http.StatusOK, toAllocationCandidatesDocument(candidates))
}

func parseRequestedResources(value string) map[string]int {
	resources := map[string]int{}
	for _, part := range strings.Split(value, ",") {
		resourceClass, amount, ok := strings.Cut(part, ":")
		if !ok {
			continue
		}

		parsedAmount, err := strconv.Atoi(amount)
		if err != nil {
			continue
		}
		resources[resourceClass] = parsedAmount
	}

	return resources
}
