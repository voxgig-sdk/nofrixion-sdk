# Nofrixion SDK utility: feature_add
module NofrixionUtilities
  FeatureAdd = ->(ctx, f) {
    ctx.client.features << f
  }
end
